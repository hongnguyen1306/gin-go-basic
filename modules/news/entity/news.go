package entity

import (
	"time"
)

type News struct {
	tableName struct{}  `pg:"news"`
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatorId string    `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewsUpdate struct {
	tableName struct{}  `pg:"news"`
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}
