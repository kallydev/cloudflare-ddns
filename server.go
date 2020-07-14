package cloudflare_ddns

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/kallydev/cloudflare-ddns/common/network"
	pb "github.com/kallydev/cloudflare-ddns/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"log"
	"net"
	"strconv"
)

type Server struct {
	context context.Context
	cancel  context.CancelFunc
	api     *cloudflare.API
	conf    *Config

	pb.CloudFlareDDNSServer
}

func NewServer(conf *Config) *Server {
	api, err := cloudflare.New(conf.CloudFlare.Key, conf.CloudFlare.Email)
	if err != nil {
		log.Panicln(err)
	}
	return &Server{
		api:  api,
		conf: conf,
	}
}

func (svr *Server) RefreshDDNSRecord(ctx context.Context, request *pb.RefreshDNSRecordRequest) (*pb.RefreshDNSRecordResponse, error) {
	if request.Key != svr.conf.CloudFlare.Key {
		return &pb.RefreshDNSRecordResponse{Message: "Key verification failed"}, nil
	}
	records, err := svr.GetDNSRecords()
	if err != nil {
		return &pb.RefreshDNSRecordResponse{Message: "Failed to get dns records"}, err
	}
	clientIP, err := svr.GetClientIP(ctx)
	if err != nil {
		return &pb.RefreshDNSRecordResponse{Message: "Error parsing client IP"}, err
	}
	recordFound := false
	for _, record := range records {
		if recordFound = record.Name == request.Name; recordFound {
			if record.Content != clientIP {
				if err := svr.UpdateDNSRecords(record.ID, record.Name, clientIP); err != nil {
					log.Println(err)
					break
				}
				log.Printf("Update %s from %s to %s", record.Name, record.Content, clientIP)
			}
			break
		}
	}
	if !recordFound {
		_type := "A"
		if network.IsIPv6(clientIP) {
			_type = "AAAA"
		}
		if err := svr.CreateDNSRecord(_type, request.Name, clientIP); err != nil {
			return &pb.RefreshDNSRecordResponse{Message: "Failed to create DNS record"}, err
		}
		log.Printf("Create a DNS record %s content is %s", request.Name, clientIP)
	}
	return &pb.RefreshDNSRecordResponse{Message: "OK"}, nil
}

func (svr *Server) GetClientIP(ctx context.Context) (ip string, err error) {
	if pr, ok := peer.FromContext(ctx); ok {
		host, _, err := net.SplitHostPort(pr.Addr.String())
		if err != nil {
			return "", err
		}
		return host, nil
	}
	return "", errors.New("")
}

func (svr *Server) GetDNSRecords() ([]cloudflare.DNSRecord, error) {
	return svr.api.DNSRecords(svr.conf.CloudFlare.ZoneID, cloudflare.DNSRecord{})
}

func (svr *Server) CreateDNSRecord(_type, name, context string) error {
	_, err := svr.api.CreateDNSRecord(svr.conf.CloudFlare.ZoneID, cloudflare.DNSRecord{
		Type:    _type,
		Name:    name,
		Content: context,
		TTL:     120,
	})
	return err
}

func (svr *Server) UpdateDNSRecords(recordID, name, context string) error {
	return svr.api.UpdateDNSRecord(svr.conf.CloudFlare.ZoneID, recordID, cloudflare.DNSRecord{
		Name:    name,
		Content: context,
		TTL:     120,
	})
}

func (svr *Server) Start() error {
	creds, err := credentials.NewServerTLSFromFile(svr.conf.TLS.Cert, svr.conf.TLS.Key)
	if err != nil {
		return err
	}
	lis, err := net.Listen("tcp", net.JoinHostPort(svr.conf.Server.Host, strconv.Itoa(svr.conf.Server.Port)))
	if err != nil {
		return err
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterCloudFlareDDNSServer(s, svr)
	return s.Serve(lis)
}
