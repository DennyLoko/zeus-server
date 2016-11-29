package model

import "database/sql/driver"

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

func (u *Gender) Scan(value interface{}) error {
	*u = Gender(value.([]byte))

	return nil
}

func (u Gender) Value() (driver.Value, error) {
	return string(u), nil
}
