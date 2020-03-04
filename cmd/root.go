package cmd

import (
	"github.com/pavkozlov/organizer/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blogserver",
	Short: "Blogserver app",
	Long:  `Personal project.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(func() {
		databaseURL := config.DbURL(config.BuildDBConfig())
		config.Db, _ = gorm.Open("postgres", databaseURL)
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
