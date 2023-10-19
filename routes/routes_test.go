package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luigiMinardi/alurachallenge-backend-7/models"
	"github.com/luigiMinardi/alurachallenge-backend-7/routes"
	"github.com/luigiMinardi/alurachallenge-backend-7/utils"
)

type ReviewInput struct {
    Name    string  `json:"name"`
    Review  string  `json:"review"`
    Image   string  `json:"image"`
}

func init() {
    utils.DB_NAME = utils.TEST_DB
}

func TestRouterAddReview(t *testing.T) {
    postReview := ReviewInput{"Fooa", "Bazz", "/9j/4AAQSkZJRgABAQAAAAAAAAD/eIZnCh0wTiTTcy2rkadlo5ICAQCB//9k="}

    data, err := json.Marshal(postReview)
    if err != nil {
        t.Fatal(err)
    }

    req, err := http.NewRequest("POST","http://url:1234/reviews", bytes.NewBuffer(data))
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    routes.Router(w, req)

    if status := w.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    var review models.Review
    json.Unmarshal(w.Body.Bytes(), &review)
}

func TestRouterGetReviewsHome(t *testing.T) {
    req, err := http.NewRequest("GET", "http://url:1234/reviews-home", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    routes.Router(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }


    var review []models.Review
    json.Unmarshal(w.Body.Bytes(), &review)
}

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

func TestRouterEditReview(t *testing.T) {
    postReview := ReviewInput{"Edited", "Review", "/9j/4AAQSkZJRgABAQAAAAAAAAD/eIZnCh0wTiTTcy2rkadlo5ICAQCB//9k="}

    data, err := json.Marshal(postReview)
    if err != nil {
        t.Fatal(err)
    }

    req, err := http.NewRequest("PUT","http://url:1234/reviews/1", bytes.NewBuffer(data))
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

func TestRouterDeleteReview(t *testing.T) {
    req, err := http.NewRequest("DELETE", "http://url:1234/reviews/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    routes.Router(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"message":"Everything worked as expected"}`
    if w.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
    }
}
