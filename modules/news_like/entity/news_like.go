package entity

import "time"

type Like struct {
	tableName struct{}  `pg:"news_like"`
	UserId    string    `json:"userId"`
	NewsId    string    `json:"newsId"`
	CreatedAt time.Time `json:"createdAt"`
}


