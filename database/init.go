package database

import (
	"test/amartha/config"
	"test/amartha/usecase/model"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Master *sqlx.DB
}

var handleSQLXConnect = sqlx.Connect

func Init(cfg config.Database) (*DB, error) {
	db, err := handleSQLXConnect("mysql", cfg.Credential)
	if err != nil {
		return nil, err
	}
	return &DB{
		Master: db,
	}, nil
}

type DBInterface interface {
	CreateShortenCode(shorten *model.ShortlnRequest) (err error)
	GetShortenByCode(code string) (shorten *model.ShortlnRequest)
	CountVisitingURL(code string) (err error)
}
