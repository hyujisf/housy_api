package housesdto

type HouseResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name" validate:"required"`
	CityId    int    `json:"city_id" form:"city_id" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Price     string `json:"price" form:"price" validate:"required"`
	TypeRent  string `json:"type_rent" form:"type_rent" validate:"required"`
	Amenities string `json:"amenities" form:"amenities" validate:"required"`
	BedRoom   string `json:"bed_room" form:"bedRoom" validate:"required"`
	BathRoom  string `json:"bath_room" form:"bathRoom" validate:"required"`
}
