package entity

import "time"

type Todo struct {
	ID        int    `gorm:"primaryKey;not null;type:int" json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
