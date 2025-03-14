package main

import (
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/database"
	"go/npcGen/npc"
	"log"
	"strings"
)

func main() {
	// Read in Database Config file
	config, err := configuration.ReadConfig("configurations/dcbonf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Run all Pre-Flight checks
	err = database.DBPreFlight(config)
	if err != nil {
		log.Fatalf("failure in DBPreFlight: %s ", err)
	}

	// Initialize database per server mode selection
	mode := strings.ToLower(config.Server.Mode)
	log.Printf("reading in config mode option %s", mode)
	switch mode {
	case "dev-db":
		err := database.InitDB(config)
		if err != nil {
			log.Fatalf("failed to init database... error: %s", err)
		}
	case "dev-csv":
		log.Print("Skipping database initialization")
	default:
		log.Fatalf("no mode matching %s. Please check configurations/dcbonf.yaml", mode)
	}

	// Create NPC
	npc_object, err := npc.CreateNPC(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----- OUTPUT -----")
	fmt.Println(npc.DataToJSON(npc_object))
	// fmt.Println(npc_object.OCEAN.Text)
}
