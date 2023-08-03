package routes

import (
	"log"
	"net/http"
)

func HandleRequest() {
    log.Fatal(http.ListenAndServe(":8000", nil))
}
