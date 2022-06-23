package models

type Tool struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
