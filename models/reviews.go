package models

import (
	"github.com/luigiMinardi/alurachallenge-backend-7/db"
)

type Review struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
    Review  string  `json:"review"`
    Image   string  `json:"image"`
}

func CreateReview(name string, review string, image string) {
    db := db.ConnectWithDB()
    err := db.Ping()
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("insert into reviews(name, review, image) values($1, $2, $3)", name, review, image)
    if err != nil {
        panic(err)
    }

    
    defer db.Close()
}
