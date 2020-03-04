package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pavkozlov/organizer/cmd"
)

var err error

func main() {
	cmd.Execute()
}
