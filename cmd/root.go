package cmd

import (
	"github.com/pavkozlov/organizer/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "organizer",
	Short: "organizer app",
	Long:  `Personal project.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
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
