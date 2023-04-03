package main

import (
	"challenges_8/server"

	_ "github.com/lib/pq"
)

func main() {
	//run server

	server.NewServer()
}
