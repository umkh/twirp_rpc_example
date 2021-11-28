package dto

import "time"

type Book struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	CreatedAt *time.Time `json:"created_at"`
}
