package entity

import (
	"app/common"
	"time"
)

type Like struct {
	tableName struct{}           `pg:"news_like"`
	NewsId    string             `json:"newsId" pg:",pk"`
	UserId    string             `json:"userId" pg:",pk"`
	CreatedAt time.Time          `json:"createdAt"`
	User      *common.SimpleUser `json:"user" pg:"rel:has-one"`
}
