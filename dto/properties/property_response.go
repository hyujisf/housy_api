package propertiesdto

import (
	"encoding/json"
)

type PropertyResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name" form:"name" validate:"required"`
	City      int             `json:"city"`
	Address   string          `json:"address" form:"address" validate:"required"`
	Price     float64         `json:"price" form:"price" validate:"required"`
	TypeRent  string          `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities json.RawMessage `json:"amenities" form:"amenities" validate:"required"`
	Bedroom   int             `json:"bedroom" form:"bedroom" validate:"required"`
	Bathroom  string          `json:"bathroom" form:"bathroom" validate:"required"`
}
