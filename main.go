package main

import (
	"github.com/Narianapereira/api-go-gin/database"
	"github.com/Narianapereira/api-go-gin/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequest()
}
