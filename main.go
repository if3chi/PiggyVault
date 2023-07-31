package main

import (
	"fmt"
	"log"

	"github.com/if3chi/PiggyVault/db"
	"github.com/if3chi/PiggyVault/routes"
)

func main() {
	fmt.Println("PiggyVault")
	db, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Init(); err != nil {
		log.Fatal(err.Error())
	}

	server := routes.NewApiServer(":3000", db)
	server.Run()
}
