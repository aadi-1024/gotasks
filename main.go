package main

import (
	"github.com/aadi-1024/gotasks/database"
	"github.com/aadi-1024/gotasks/cmd"
)

func main() {
	database.Setup()
	defer database.Db.Close()
	cmd.Execute()
}