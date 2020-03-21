package cmd

import (
	"github.com/pavkozlov/organizer/urls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		router := urls.SetupRouter()
		router.Run()
	},
}
