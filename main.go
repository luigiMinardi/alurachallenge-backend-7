package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    fmt.Println("Starting Service")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
