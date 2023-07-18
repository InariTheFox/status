/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/inarithefox/status/api"
	"github.com/inarithefox/status/app"
	"github.com/inarithefox/status/web"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the status server",
	RunE: func(cmd *cobra.Command, args []string) error {
		interrupt := make(chan os.Signal, 1)

		return runServer(interrupt)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runServer(interrupt chan os.Signal) error {

	server, err := app.NewServer()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer server.Shutdown()
	defer func() {

	}()

	if err = server.Start(); err != nil {
		log.Fatalln(err)
		return err
	}

	_, err = api.Init(server)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	web.New(server)

	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	log.Println("shutdown command received")

	return nil
}
