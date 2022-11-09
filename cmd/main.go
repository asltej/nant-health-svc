package main

import (
	"nant-health-svc/db"
	"nant-health-svc/handlers"
)

func main() {
	// Initialize DB client
	db.NewClient()

	// Initialize APIs
	handlers.Init()
}
