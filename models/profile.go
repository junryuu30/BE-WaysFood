package models

import "time"

type Profile struct {
	ID        int       `json:"id"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	UserID    int       `json:"userId"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

type ProfileResponse struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
