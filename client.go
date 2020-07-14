package cloudflare_ddns

import (
	"context"
	pb "github.com/kallydev/cloudflare-ddns/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

type Client struct {
	context context.Context
	cancel  context.CancelFunc
	conf    *Config
	client  pb.CloudFlareDDNSClient
}

func NewClient(conf *Config) *Client {
	ctx, cancel := context.WithCancel(context.Background())
	return &Client{
		context: ctx,
		cancel:  cancel,
		conf:    conf,
	}
}

func (cli *Client) Start() (err error) {
	var conn *grpc.ClientConn
	if cli.conf.TLS != nil {
		creds, err := credentials.NewClientTLSFromFile(cli.conf.TLS.Cert, "")
		if err != nil {
			return err
		}
		conn, err = grpc.Dial(cli.conf.Client.Server, grpc.WithTransportCredentials(creds))
	} else {
		conn, err = grpc.Dial(cli.conf.Client.Server, grpc.WithInsecure(), grpc.WithBlock())
	}
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli.client = pb.NewCloudFlareDDNSClient(conn)
	for {
		select {
		case <-cli.context.Done():
		default:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			r, err := cli.client.RefreshDDNSRecord(ctx, &pb.RefreshDNSRecordRequest{
				Name: cli.conf.Client.Domain,
				Key:  cli.conf.CloudFlare.Key,
			})
			if err != nil {
				log.Println(err)
			}
			if r != nil && r.Message != "OK" {
				log.Println(r.Message)
			}
			cancel()
		}
		time.Sleep(time.Second * 5)
	}
}
