package housesdto

import "encoding/json"

type AddHouseRequest struct {
	Name      string          `json:"name" form:"name" validate:"required"`
	CityId    int             `json:"city_id" form:"city_id" validate:"required"`
	Address   string          `json:"address" form:"address" validate:"required"`
	Price     float64         `json:"price" form:"price" validate:"required"`
	TypeRent  string          `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities json.RawMessage `json:"amenities" form:"amenities" validate:"required"`
	Bedroom   int             `json:"bedroom" form:"bedroom" validate:"required"`
	Bathroom  int             `json:"bathroom" form:"bathroom" validate:"required"`
}

type UpdateHouseRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	CityId    string `json:"city_id" form:"city_id" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Price     string `json:"price" form:"price" validate:"required"`
	TypeRent  string `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities string `json:"amenities" form:"amenities" validate:"required"`
	Bedroom   string `json:"bedroom" form:"bedroom" validate:"required"`
	Bathroom  string `json:"bathroom" form:"bathroom" validate:"required"`
}
