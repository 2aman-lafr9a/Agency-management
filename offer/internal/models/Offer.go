package models

import _ "gorm.io/gorm"

type Offer struct {
	ID          string  `gorm:"primaryKey;type:varchar(255);not null;unique"`
	Name        string  `gorm:"type:varchar(255);not null;unique"`
	Description string  `gorm:"type:varchar(255);not null;unique"`
	Price       float64 `gorm:"type:float;not null;unique"`
	Date        string  `gorm:"type:varchar(255);not null;unique"`
	AgencyID    string  `gorm:"type:varchar(255);not null;unique;references:agencies"`
}

func (o Offer) Error() string {
	err := o.Error()
	return err
}
