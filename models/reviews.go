package models

import "github.com/luigiMinardi/alurachallenge-backend-7/db"

type Review struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
    Review  string  `json:"review"`
    Image   []byte  `json:"image"`
}

func CreateReview(name, review string, image []byte) {
    db := db.ConnectWithDB()

    sqlToInsert, err := db.Prepare("insert into reviews(name, review, image) values($1, $2, $3)")
    if err != nil {
        panic(err.Error())
    }

    sqlToInsert.Exec(name, review, image)
    defer db.Close()
}
