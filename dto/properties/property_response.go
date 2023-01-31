package propertiesdto

import "gorm.io/datatypes"

type PropertyResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name" form:"name" validate:"required"`
	City      int            `json:"city"`
	Address   string         `json:"address" form:"address" validate:"required"`
	Price     float64        `json:"price" form:"price" validate:"required"`
	TypeRent  string         `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities datatypes.JSON `json:"amenities" form:"amenities" validate:"required"`
	Bedroom   int            `json:"bedroom" form:"bedroom" validate:"required"`
	Bathroom  int            `json:"bathroom" form:"bathroom" validate:"required"`
	// Image     datatypes.JSON `json:"image" form:"image" validate:"required"`
}
