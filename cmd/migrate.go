package cmd

import (
	"github.com/pavkozlov/organizer/applications/account"
	"github.com/pavkozlov/organizer/applications/todo"
	"github.com/pavkozlov/organizer/organizer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		organizer.Db.AutoMigrate(
			&todo.Todo{},
			&account.User{},
			&account.Sessions{},
		)
	},
}
