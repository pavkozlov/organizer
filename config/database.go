package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		DBName:   "golang",
		Password: "01234567",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}
