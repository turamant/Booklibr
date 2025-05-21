package data

import (
    "time"
)

// Добавляем структурные теги, чтобы управлять тем, как будут выглядеть ключи в JSON.
type Book struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Title     string    `json:"title"`
    Year      int32     `json:"year"`
    Pages     int32     `json:"pages"`
    Genres    []string  `json:"genres"`
    Author    string    `json:"author"`
}