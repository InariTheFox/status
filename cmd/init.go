package cmd

import (
	"github.com/inarithefox/status/database"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initalize the database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.Connect("postgresql://status_svc:status@localhost/status?sslmode=disable")
		if err != nil {
			return
		}

		db.CreateTables()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
