package entities

import "time"

type Book struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Author    string     `json:"author" db:"author"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}
