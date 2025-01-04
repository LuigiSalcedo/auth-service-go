package models

type User struct {
	ID       uint   `grom:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}
