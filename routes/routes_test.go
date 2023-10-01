package routes_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luigiMinardi/alurachallenge-backend-7/models"
	"github.com/luigiMinardi/alurachallenge-backend-7/routes"
)

func TestRouterGetReview(t *testing.T) {
    req, err := http.NewRequest("GET", "http://url:1234/reviews/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    routes.Router(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }


    var review models.Review
    json.Unmarshal(w.Body.Bytes(), &review)
}
