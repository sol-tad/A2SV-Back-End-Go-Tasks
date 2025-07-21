package main

import (
	"task_manager/config"
	"task_manager/router"
)

func main() {
	config.ConnectDB()
	r := router.SetupRouter()
	r.Run(":8080")
}