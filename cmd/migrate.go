package cmd

import (
	"github.com/pavkozlov/organizer/applications/account"
	"github.com/pavkozlov/organizer/models"
	"github.com/pavkozlov/organizer/settings"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		settings.Db.AutoMigrate(&models.Todo{}, &account.User{})
	},
}
