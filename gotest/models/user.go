package models

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}
