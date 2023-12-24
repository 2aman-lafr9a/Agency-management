package models

import "gorm.io/gorm"

type Agency struct {
	gorm.Model

	ID   string `gorm:"primaryKey;type:varchar(255);not null;unique;autoIncrement:true"`
	Name string `gorm:"type:varchar(255);not null;unique"`
}

func (a Agency) Error() string {
	err := a.Error()
	return err
}
