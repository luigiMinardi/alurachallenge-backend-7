package controllers

import (
    //"encoding/base64"
    "encoding/json"
    "log"
    "net/http"
    "github.com/luigiMinardi/alurachallenge-backend-7/models"
)

func AddReview(w http.ResponseWriter, r *http.Request) {
    var data models.Review
    json.NewDecoder(r.Body).Decode(&data)
    /*
    decodedImage, err := base64.StdEncoding.DecodeString(data.Image)
    if err != nil {
        panic(err)
    }
    */
    models.CreateReview(data.Name, data.Review, data.Image)

    w.WriteHeader(http.StatusOK)

    resp := make(map[string]string)
    resp["message"] = "Everything worked as expected"
    response, err := json.Marshal(resp)
    if err != nil {
        log.Panic(err)
    }
    w.Write(response)
}
