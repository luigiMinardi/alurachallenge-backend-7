package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"strconv"

	"github.com/luigiMinardi/alurachallenge-backend-7/models"
)

func AddReview(w http.ResponseWriter, r *http.Request) {
    var data models.Review
    json.NewDecoder(r.Body).Decode(&data)
    models.CreateReview(data.Name, data.Review, data.Image)

    w.WriteHeader(http.StatusCreated)

    resp := make(map[string]string)
    resp["message"] = "Everything worked as expected"
    response, err := json.Marshal(resp)
    if err != nil {
        panic(err)
    }
    w.Write(response)
}

func EditReview(w http.ResponseWriter, r *http.Request) {
    path := strings.Split(r.URL.Path, "/")[1:]
    id := path[len(path)-1]

    idToInt, err := strconv.Atoi(id)
    if err != nil {
        panic(err)
    }

    var data models.Review
    json.NewDecoder(r.Body).Decode(&data)
    data.Id = idToInt

    models.UpdateReview(idToInt, data.Name, data.Review, data.Image)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}
