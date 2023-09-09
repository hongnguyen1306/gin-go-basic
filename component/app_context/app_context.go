package app_context

import (
	"app/common"

	"github.com/go-pg/pg/v10"
)

type AppContext interface {
	SecretKey() string
	GetMainDBConnection() *pg.DB
}

type appCtx struct {
	db        *pg.DB
	secretKey string
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetMainDBConnection() *pg.DB {
	return ctx.db
}

func NewAppContext(db *pg.DB, config *common.Config) *appCtx {
	return &appCtx{db, config.SecretKey}
}
