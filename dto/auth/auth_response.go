package authdto

type LoginResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
