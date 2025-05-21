package data

import (
    "time"
)

// Добавляем структурные теги, чтобы управлять тем, как будут выглядеть ключи в JSON.
type Movie struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Title     string    `json:"title"`
    Year      int32     `json:"year"`
    Runtime   int32     `json:"runtime"`
    Genres    []string  `json:"genres"`
    Version   int32     `json:"version"`
}