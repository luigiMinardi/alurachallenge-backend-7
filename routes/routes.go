package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/luigiMinardi/alurachallenge-backend-7/controllers"
	"github.com/luigiMinardi/alurachallenge-backend-7/middleware"
	"github.com/luigiMinardi/alurachallenge-backend-7/utils"
)


func HandleRequest() {
    utils.DB_NAME = utils.DEFAULT_DB
    mux := http.NewServeMux()
    mux.HandleFunc("/", Router)
    log.Fatal(http.ListenAndServe(":8000", middleware.CORSMiddleware(middleware.ContentTypeMiddleware(mux))))
}

func Router(w http.ResponseWriter, r *http.Request) {
    p := strings.Split(r.URL.Path, "/")[1:]
    log.Println("p: ",p)
    log.Println("method: ", r.Method)

    var h http.HandlerFunc
    switch {
    case p[0] == "reviews-home" && r.Method == "GET":
        log.Println("reviews home - GET")
        h = controllers.GetReviewsHome
    case p[0] == "reviews" && r.Method == "POST":
        log.Println("reviews - POST")
        h = controllers.AddReview
    case len(p) >= 2 && p[0] == "reviews" && p[1] != "" && r.Method == "GET":
        log.Println("reviews - GET")
        h = controllers.GetReview
    case len(p) >= 2 && p[0] == "reviews" && p[1] != "" && r.Method == "PUT":
        log.Println("reviews - PUT")
        h = controllers.EditReview
    case len(p) >= 2 && p[0] == "reviews" && p[1] != "" && r.Method == "DELETE":
        h = controllers.DeleteReview
    default:
        http.NotFound(w,r)
        return
    }
    h.ServeHTTP(w,r)
}
