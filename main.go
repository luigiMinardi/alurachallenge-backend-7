package main

import (
	"fmt"

	"github.com/luigiMinardi/alurachallenge-backend-7/routes"
)

func main() {
    fmt.Println("Starting Service")
    routes.HandleRequest()
}
