package main

import "nexus/server"

func main() {
	err := server.StartServer()
	if err != nil {
		return
	}
}
