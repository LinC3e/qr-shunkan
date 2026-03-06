package main

import "github.com/LinC3e/shunkan-qr/internal/router"

func main() {

	r := router.Setup()

	r.Run(":8080")

}