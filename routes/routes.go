package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/luigiMinardi/alurachallenge-backend-7/middleware"
)

func HandleRequest() {
    r := http.NewServeMux()
    r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
       io.WriteString(rw, `{"message": "hello world.."}`)
    })
    
    log.Fatal(http.ListenAndServe(":8000", middleware.CORSMiddleware(middleware.ContentTypeMiddleware(r))))
}
