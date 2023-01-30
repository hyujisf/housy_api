package models

import (
	"encoding/json"
	"time"
)

type House struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	Name      string          `json:"name" gorm:"type: varchar(255)"`
	City      CityResponse    `json:"city" gorm:"foreignKey:CityId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CityId    int             `json:"city_id"`
	Address   string          `json:"address" gorm:"type: text"`
	Price     float64         `json:"price" gorm:"type: decimal(10,2)"`
	TypeRent  string          `json:"type_rent" gorm:"type: varchar(255)"`
	Amenities json.RawMessage `json:"amenities" gorm:"type:json"`
	BedRoom   int             `json:"bed_room" gorm:"type: int"`
	BathRoom  int             `json:"bath_room" gorm:"type: int"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type HouseResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	CityId    int             `json:"city_id"`
	Address   string          `json:"address"`
	Price     float64         `json:"price"`
	TypeRent  string          `json:"type_rent"`
	Amenities json.RawMessage `json:"amenities"`
	BedRoom   int             `json:"bed_room"`
	BathRoom  int             `json:"bath_room"`
}

func (HouseResponse) TableName() string {
	return "users"
}
