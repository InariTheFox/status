/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/inarithefox/status/database"
	"github.com/spf13/cobra"
)

// dropCmd represents the drop command
var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop all database tables",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("confirm").Value.String() == "true" {
			db, err := database.Connect("postgresql://status_svc:status@localhost/status?sslmode=disable")
			if err != nil {
				return
			}

			db.DropTables()
		} else {
			log.Fatal("Drop tables aborted as \"confirm\" flag was not set.")
		}
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
	dropCmd.Flags().BoolP("confirm", "", false, "Confirm must be supplied in order to drop all db tables.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dropCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dropCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
