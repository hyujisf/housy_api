package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        int            `json:"id" gorm:"primary_key:auto_increment"`
	Fullname  string         `json:"fullname" gorm:"type: varchar(255)"`
	Email     string         `json:"email" gorm:"type: varchar(255)"`
	Username  string         `json:"username" gorm:"type: varchar(255)"`
	Password  string         `json:"-" gorm:"type: varchar(255)"`
	ListAs    ListAsResponse `json:"listAs" gorm:"foreignKey:ListAsId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ListAsId  int            `json:"list_as_id"`
	Gender    string         `json:"gender" gorm:"type: varchar(255)"`
	Address   string         `json:"address" gorm:"type: text"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (UsersProfileResponse) TableName() string {
	fmt.Println("Users Tables is created")
	return "users"
}
