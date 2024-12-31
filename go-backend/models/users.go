package models

type User struct {
	Id   string `json:"id" gorm:"primaryKey;type:uuid;"`
	Name string `json:"name"`
}
