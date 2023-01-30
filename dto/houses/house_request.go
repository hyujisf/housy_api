package housesdto

import "encoding/json"

type AddHouseRequest struct {
	Name      string          `json:"name" form:"name" validate:"required"`
	CityId    int             `json:"city_id" form:"city_id" validate:"required"`
	Address   string          `json:"address" form:"address" validate:"required"`
	Price     float64         `json:"price" form:"price" validate:"required"`
	TypeRent  string          `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities json.RawMessage `json:"amenities" form:"amenities" validate:"required"`
	BedRoom   int             `json:"bed_room" form:"bedRoom" validate:"required"`
	BathRoom  int             `json:"bath_room" form:"bathRoom" validate:"required"`
}

type UpdateHouseRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	CityId    string `json:"city_id" form:"city_id" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Price     string `json:"price" form:"price" validate:"required"`
	TypeRent  string `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities string `json:"amenities" form:"amenities" validate:"required"`
	BedRoom   string `json:"bed_room" form:"bedRoom" validate:"required"`
	BathRoom  string `json:"bath_room" form:"bathRoom" validate:"required"`
}
