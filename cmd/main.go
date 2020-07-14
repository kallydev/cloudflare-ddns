package main

import (
	"github.com/kallydev/cloudflare-ddns"
	"github.com/spf13/cobra"
	"log"
)

var (
	configPath string
)

var rootCmd = &cobra.Command{
	Use:     "cloudflare-ddns",
	Version: "1.0.0",
	Long:    "A CloudFlare DDNS service",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Long:  `CloudFlare DDNS server`,
	Run: func(cmd *cobra.Command, args []string) {
		svr := cloudflare_ddns.NewServer(cloudflare_ddns.NewConfig(configPath))
		if err := svr.Start(); err != nil {
			log.Panicln(err)
		}
	},
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Long:  `CloudFlare DDNS client`,
	Run: func(cmd *cobra.Command, args []string) {
		cli := cloudflare_ddns.NewClient(cloudflare_ddns.NewConfig(configPath))
		if err := cli.Start(); err != nil {
			log.Panicln(err)
		}
	},
}

func init() {
	rootPf := rootCmd.PersistentFlags()
	rootPf.StringVarP(&configPath, "config", "c", "config.json", "config file path")

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Panicln(err)
	}
}
