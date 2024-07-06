package main

import (
	"fmt"
	"nexus/cmd/server"
	"nexus/pkg/db"
)

func main() {
	dbPool, dbErr := db.InitDB()
	if dbErr != nil {
		return
	}
	db.Connection = dbPool
	defer fmt.Println("Closing database connection")
	defer dbPool.Close()
	err := server.StartServer()
	if err != nil {
		return
	}
}
