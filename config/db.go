package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:password@(db:3306)/gantt_taro")
	if err != nil {
		return nil, err
	}

	return db, nil
}
