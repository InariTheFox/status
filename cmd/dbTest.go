/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/inarithefox/status/database"
	"github.com/spf13/cobra"
)

// dbTestCmd represents the dbTest command
var dbTestCmd = &cobra.Command{
	Use:   "dbTest",
	Short: "Test the database connection",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.Connect("postgresql://status_svc:status@localhost/status?sslmode=disable")
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Test connection succeeded.")
	},
}

func init() {
	rootCmd.AddCommand(dbTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
