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

func ShowReview(id int) Review {
    db := db.ConnectWithDB()
    
    query, err := db.Query("select * from reviews where id=$1", id)
    if err != nil {
        panic(err)
    }
    var review Review

    query.Next()
    err2 := query.Scan(&review.Id, &review.Name, &review.Review, &review.Image)
    if err2 != nil {
        panic(err)
    }

    defer db.Close()
    return review
}

func RemoveReview(id int) {
    db := db.ConnectWithDB()

    _, err := db.Exec("delete from reviews where id=$1", id)
    if err != nil {
        panic(err)
    }

    defer db.Close()
}

func GetThreeRandomReviews() ([]Review, error) {
    db := db.ConnectWithDB()

    query, err := db.Query("select * from reviews order by random() limit 3")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    var reviews []Review

    for query.Next() {
        var review Review
        if err := query.Scan(&review.Id, &review.Name,
            &review.Review, &review.Image); err != nil {
                return reviews, err
            }
        reviews = append(reviews, review)
    }

    if err = query.Err(); err != nil {
        return reviews, err
    }
    return reviews, nil
}
