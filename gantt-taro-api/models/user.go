package models

type User struct {
	Id       int    `json:"id" db:"id"`
	UniqueId string `json:"unique_id" db:"unique_id"`
}
