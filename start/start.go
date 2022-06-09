package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"guide/ent"
)

func main() {
	config := LoadConfig("config.json")
	client, err := ent.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", config.DriverName, err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
