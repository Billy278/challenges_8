package model

import "time"

type Book struct {
	Id        uint64     `json:"id" bson:"id"`
	Title     string     `json:"title" bson:"title"`
	Author    string     `json:"author" bson:"author"`
	Desc      string     `json:"desc" bson:"des"`
	Create_at *time.Time `json:"create_at" bson:"create_at"`
	Update_at *time.Time `json:"update_at" bson:"update_at"`
	Delete_at *time.Time `json:"delete_at" bson:"delete_at"`
}
