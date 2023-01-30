package authdto

type RegisterRequest struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"Email" validate:"required"`
	Username string `gorm:"type: varchar(255)" json:"username" validate:"required"`
	ListAsId int    `gorm:"type: varchar(255)" json:"list_as_id" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Address  string `gorm:"type: varchar(255)" json:"address" validate:"required"`
}
type LoginRequest struct {
	Username string `gorm:"type: varchar(255)" json:"username" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
