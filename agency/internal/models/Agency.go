package models

import (
	"gorm.io/gorm"
)

type Agency struct {
	gorm.Model

	ID          string  `gorm:"primaryKey;type:varchar(255);not null;unique;autoIncrement:true"`
	Name        string  `gorm:"type:varchar(255);not null;unique"`
	Description string  `gorm:"type:varchar(255);not null"`
	Plan        string  `gorm:"type:varchar(255);not null"`
	Offers      []Offer `gorm:"foreignKey:AgencyID"`
}

type Offer struct {
	ID          string    `gorm:"primaryKey;type:varchar(255);not null;unique;autoIncrement:true"`
	Name        string    `gorm:"type:varchar(255);not null;unique"`
	Description string    `gorm:"type:varchar(255);not null;unique"`
	Price       float64   `gorm:"type:float;not null;unique"`
	Date        string    `gorm:"type:varchar(255);not null;unique"`
	AgencyID    string    `gorm:"type:varchar(255);not null;unique;references:agencies"`
	Rating      int32     `gorm:"type:int;not null;default:0"`
	OfferType   OfferType `gorm:"type:int;not null;default:0"`
}

type OfferType int

const (
	Normal OfferType = iota
	Special
	Important
	Ultimate
)

func (a Agency) Error() string {
	err := a.Error()
	return err
}
