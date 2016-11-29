package model

type Account struct {
	ID       uint   `gorm:"primary_key"`
	User     string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Sex      string `gorm:"type:char(1)"`
	Email    string `gorm:"unique"`
	Slots    uint
}
