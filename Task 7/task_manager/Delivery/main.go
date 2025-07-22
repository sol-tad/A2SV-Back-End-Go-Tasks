package main

import (
	"task_manager/Delivery/router"
	"task_manager/config"
)

func main() {
	config.ConnectDB()
	r := router.SetupRouter()
	r.Run(":8080")
}