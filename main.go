package main

import (
	"leaffs/cmd"
	"leaffs/db"
	"log"
)

func main() {
	cfg := db.Config{DatabasePath: "C:/Users/flufs/Desktop/db/db.db"}
	store, err := db.NewSQLStore(cfg)
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}
	defer store.Close()
	cmd.CreateServer()
}
