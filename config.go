package cloudflare_ddns

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	CloudFlare *ConfigCloudFlare `json:"cloud_flare"`
	TLS        *ConfigTLS        `json:"tls"`
	Client     *ConfigClient     `json:"client"`
	Server     *ConfigServer     `json:"server"`
}

type ConfigCloudFlare struct {
	Email     string `json:"email"`
	Key       string `json:"key"`
	ZoneID    string `json:"zone_id"`
	AccountID string `json:"account_id"`
}

type ConfigTLS struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type ConfigClient struct {
	Server string `json:"server"`
	Domain string `json:"domain"`
}

type ConfigServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewConfig(filename string) *Config {
	conf := &Config{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicln(err)
	}
	if err := conf.Unmarshal(data); err != nil {
		log.Panicln(err)
	}
	return conf
}

func (conf *Config) Marshal() ([]byte, error) {
	return json.Marshal(conf)
}

func (conf *Config) Unmarshal(data []byte) error {
	return json.Unmarshal(data, conf)
}
