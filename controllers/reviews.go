package controllers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	//"github.com/luigiMinardi/alurachallenge-backend-7/models"
)

type Image struct {
    MimeType        string `json:"mime"`
    Base64Encoded   string `json:"base64"`
}
type ReviewJSON struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
    Review  string  `json:"review"`
    Image   string  `json:"image"`
}

func AddReview(w http.ResponseWriter, r *http.Request) {
    var data ReviewJSON
    json.NewDecoder(r.Body).Decode(&data)
    //var serializedData models.Review

    decodedImage, err := base64.StdEncoding.DecodeString(data.Image)
    if err != nil {
        panic(err)
    }
    log.Println(decodedImage)
    //models.CreateReview(data.Name, data.Review, data.Image)
}
