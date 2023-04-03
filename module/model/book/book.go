package model

import "time"

type Book struct {
	Id        uint64     `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Desc      string     `json:"desc"`
	Create_at *time.Time `json:"create_at"`
	Update_at *time.Time `json:"update_at"`
	Delete_at *time.Time `json:"delete_at"`
}
