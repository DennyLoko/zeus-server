package model

type Party struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Exp      uint
	LeaderID uint   `gorm:"index"`
	Leader   Char   `gorm:"ForeignKey:LeaderID;AssociationForeignKey:Refer"`
	Members  []Char `gorm:"ForeignKey:PartyID"`
}
