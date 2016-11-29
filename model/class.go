package model

type Class struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
