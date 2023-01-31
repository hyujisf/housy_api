package models

import (
	"time"

	"gorm.io/datatypes"
)

type Property struct {
	ID        int            `json:"id" gorm:"primary_key:auto_increment"`
	Name      string         `json:"name" gorm:"type: varchar(255)"`
	CityId    int            `json:"city_id"`
	City      CityResponse   `json:"city" gorm:"foreignKey:city_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address   string         `json:"address" gorm:"type: text"`
	Price     float64        `json:"price" gorm:"type: decimal(10,2)"`
	TypeRent  string         `json:"type_rent" gorm:"type: varchar(255)"`
	Amenities datatypes.JSON `json:"amenities" gorm:"type: json"`
	Bedroom   int            `json:"bedroom" gorm:"type: int"`
	Bathroom  int            `json:"bathroom" gorm:"type: int"`
	Image     string         `json:"image" gorm:"type: json"`
	UserID    int            `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

type PropertyResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CityId    int            `json:"city_id"`
	City      CityResponse   `json:"city" gorm:"foreignKey:city_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address   string         `json:"address"`
	Price     float64        `json:"price"`
	TypeRent  string         `json:"type_rent"`
	Amenities datatypes.JSON `json:"amenities"`
	Bedroom   int            `json:"bedroom"`
	Bathroom  int            `json:"bathroom"`
	Image     string         `json:"image" gorm:"type: json"`
	UserID    int            `json:"-"`
}

func (PropertyResponse) TableName() string {
	return "properties"
}
