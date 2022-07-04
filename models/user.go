package models

type User struct {
	AuditTrails
	Username string `gorm:"not null;uniqueIndex" json:"username" valid:"required~ Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your full email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your full password is required, minstringlength(6)~Password as to have a minimum length of 6 chaarcters"`
	Age      int    `gorm:"not null" json:"age" valid:"required~Your Age is required"`
}
