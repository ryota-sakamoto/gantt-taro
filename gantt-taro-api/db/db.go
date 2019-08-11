package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"fmt"
)

func NewDB(config *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:3306)/gantt_taro", config.User, config.Password, config.Host))

	return db, err
}
