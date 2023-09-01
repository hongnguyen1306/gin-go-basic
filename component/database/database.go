package database

import (
	"app/modules/user/entity"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func ConnectDatabase() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "123",
		Database: "postgres",
	})

	err := createSchema(db)
	if err != nil {
		log.Println("2777", err)
		panic(err)
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
