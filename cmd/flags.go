package cmd

import (
	"github.com/inarithefox/status/utils"
	"github.com/spf13/cobra"
)

var (
	ipAddress  string
	configFile string
	port       int
)

func ParseFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&ipAddress, "ip", "", "0.0.0.0", "ip address to bind the server to")
	utils.Params.BindPFlag("ip", cmd.PersistentFlags().Lookup("ip"))

	cmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	utils.Params.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", utils.Directory+"/config.yml", "path to config.yml file")
	utils.Params.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
}
