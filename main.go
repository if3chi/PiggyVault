package main

import (
	"fmt"

	"github.com/if3chi/PiggyVault/routes"
)

func main() {
	fmt.Println("PiggyVault")
	server := routes.NewApiServer(":3000")

	server.Run()
}
