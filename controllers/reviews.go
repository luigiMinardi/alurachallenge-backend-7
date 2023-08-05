package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/luigiMinardi/alurachallenge-backend-7/models"
)

func AddReview(w http.ResponseWriter, r *http.Request) {
    var data models.Review
    json.NewDecoder(r.Body).Decode(&data)
    models.CreateReview(data.Name, data.Review, data.Image)

    w.WriteHeader(http.StatusOK)

    resp := make(map[string]string)
    resp["message"] = "Everything worked as expected"
    response, err := json.Marshal(resp)
    if err != nil {
        panic(err)
    }
    w.Write(response)
}

func EditReview(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    idToInt, err := strconv.Atoi(id)
    if err != nil {
        panic(err)
    }

    var data models.Review
    json.NewDecoder(r.Body).Decode(&data)

    models.UpdateReview(idToInt, data.Name, data.Review, data.Image)
}
