package main

import (
	"github.com/tarcea/go-crud-01/initializers"
	"github.com/tarcea/go-crud-01/models"
)

func init() {
	initializers.ConnectToDB()

}

func main() {
	initializers.Db.AutoMigrate(&models.Post{})

}
