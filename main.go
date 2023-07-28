package main

import "fmt"

func main() {
	fmt.Println("PiggyVault")
	server := NewApiServer(":3000")

	server.Run()
}
