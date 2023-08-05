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

func CreateReview(name, review string, image string) {
    db := db.ConnectWithDB()

    _, err := db.Exec("insert into reviews(name, review, image) values($1, $2, $3)", name, review, image)
    if err != nil {
        panic(err)
    }

    defer db.Close()
}

func UpdateReview(id int, name, review string, image string) {
    db := db.ConnectWithDB()

    _, err := db.Exec("update reviews set name=$1, review=$2, image=$3 where id=$4", name, review, image, id)
    if err != nil {
        panic(err)
    }

    defer db.Close()
}
