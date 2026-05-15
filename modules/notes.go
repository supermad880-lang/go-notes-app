package modules

import "time"

type Notes struct {
	ID        int       `gorm:"primary key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
