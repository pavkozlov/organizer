package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pavkozlov/organizer/cmd"
)

func main() {
	cmd.Execute()
}
