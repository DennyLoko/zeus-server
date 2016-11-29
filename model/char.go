package model

type Char struct {
	ID        uint `gorm:"primary_key"`
	Account   Account
	AccountID uint `gorm:"index"`
	Slot      uint
	Nick      string `gorm:"type:varchar(30);no null;unique"`
	Gender    Gender
	Class     Class
	ClassID   uint
	PartyID   uint
	BaseLevel uint
	JobLevel  uint
	BaseExp   uint
	JobExp    uint
	Coin      uint
	HP        uint
	MaxHP     uint
	SP        uint
	MaxSP     uint
	Str       uint
	Agi       uint
	Vit       uint
	Int       uint
	Dex       uint
	Luk       uint
}
