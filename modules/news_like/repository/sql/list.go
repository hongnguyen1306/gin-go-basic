package sql

import (
	"app/common"
	"app/modules/news_like/entity"
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/go-pg/pg/v10"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlRepo) GetUsersLikeNews(context context.Context, condition map[string]interface{}, filter *entity.Filter, paging *common.Paging) ([]common.SimpleUser, error) {
	var result []entity.Like
	if err := s.db.Model(&result).Select(); err != nil {
		return nil, err
	}

	if v := filter; v != nil {
		if v.NewsId != "" {
			if err := s.db.Model(&result).Where("news_id = ?", v.NewsId).Select(); err != nil {
				return nil, err
			}
		}
	}

	count, err := s.db.Model(&result).Count()
	if err != nil {
		return nil, err
	}
	paging.Total = count

	err = s.db.Model(&result).Relation("User").Select()
	if err != nil {
		return nil, err
	}

	var ormResult *pg.Query
	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, err
		}
		ormResult = s.db.Model(&result).Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		ormResult = s.db.Model(&result).Offset((paging.Page - 1) * paging.Limit)
	}

	if err := ormResult.Limit(paging.Limit).Order("created_at desc").Select(); err != nil {
		return nil, err
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = time.Time{}
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil

}
