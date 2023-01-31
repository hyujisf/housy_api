package propertiesdto

import "encoding/json"

type AddPropertyRequest struct {
	Name      string          `json:"name" form:"name" validate:"required"`
	CityId    int             `json:"city_id" form:"city_id" validate:"required"`
	Address   string          `json:"address" form:"address" validate:"required"`
	Price     float64         `json:"price" form:"price" validate:"required"`
	TypeRent  string          `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities json.RawMessage `json:"amenities" form:"amenities" validate:"required"`
	Bedroom   int             `json:"bedroom" form:"bedroom" validate:"required"`
	Bathroom  int             `json:"bathroom" form:"bathroom" validate:"required"`
}

type UpdatePropertyRequest struct {
	Name      string          `json:"name" form:"name" validate:"required"`
	CityId    int             `json:"city_id" form:"city_id" validate:"required"`
	Address   string          `json:"address" form:"address" validate:"required"`
	Price     float64         `json:"price" form:"price" validate:"required"`
	TypeRent  string          `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities json.RawMessage `json:"amenities" form:"amenities" validate:"required"`
	Bedroom   int             `json:"bedroom" form:"bedroom" validate:"required"`
	Bathroom  int             `json:"bathroom" form:"bathroom" validate:"required"`
}
