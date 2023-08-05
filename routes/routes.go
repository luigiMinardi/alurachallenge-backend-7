package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/luigiMinardi/alurachallenge-backend-7/controllers"
	"github.com/luigiMinardi/alurachallenge-backend-7/middleware"
)

func HandleRequest() {
    r := http.NewServeMux()
    r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
       io.WriteString(rw, `{"message": "hello world.."}`)
    })
    r.HandleFunc("/reviews", controllers.AddReview)
    r.Handle("/reviews/", http.StripPrefix("/reviews/", http.HandlerFunc(controllers.EditReview)))
    log.Fatal(http.ListenAndServe(":8000", middleware.CORSMiddleware(middleware.ContentTypeMiddleware(r))))
}
