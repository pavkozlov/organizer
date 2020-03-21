package cmd

import (
	"github.com/pavkozlov/organizer/applications/account"
	"github.com/pavkozlov/organizer/applications/todo"
	"github.com/pavkozlov/organizer/organizer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		router := organizer.SetupRouter()
		account.SetupRouter(router)
		todo.SetupRouter(router)
		router.Run(":8080")
	},
}
