package cmd

import (
	"github.com/pavkozlov/organizer/config"
	"github.com/pavkozlov/organizer/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Make a DB migration",
	Run: func(cmd *cobra.Command, args []string) {
		config.Db.AutoMigrate(&models.Todo{})
	},
}
