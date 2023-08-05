package controllers

import (
    "encoding/json"
    "net/http"
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
