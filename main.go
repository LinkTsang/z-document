package main

import (
	"fmt"
	"znote-server-go/routers"
)

// @title ZNote Server
// @version 1.0
// @description This is the ZNote server.

// @host localhost:8000
// @BasePath /api

func main() {
	r := routers.InitRouter()
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err.Error())
	}
}
