package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pavkozlov/organizer/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
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

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		credentials := viper.GetStringMapString("service.db")
		databaseURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			credentials["host"],
			credentials["port"],
			credentials["username"],
			credentials["dbname"],
			credentials["password"])

		conn, err := gorm.Open(credentials["engine"], databaseURI)
		if err != nil {
			logrus.Fatal("failed to connect database:", err)
		}
		config.Db = conn
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
