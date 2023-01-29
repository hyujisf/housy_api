package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	// Profile   ProfileResponse       `json:"profile"`
	// Products  []ProductUserResponse `json:"products"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
