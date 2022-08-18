package domain

import "time"

type Book struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Books []Book
