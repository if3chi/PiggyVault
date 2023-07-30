package main

import (
	"fmt"
	"log"

	"github.com/if3chi/PiggyVault/db"
	// "github.com/if3chi/PiggyVault/routes"
)

func main() {
	fmt.Println("PiggyVault")
	db, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%+v\n", db)

	// server := routes.NewApiServer(":3000", db)
	// server.Run()
}
