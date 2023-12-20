package models

import "gorm.io/gorm"

type Agency struct {
	gorm.Model

	Name string `gorm:"type:varchar(255);not null;unique"`
}

func (a Agency) Error() string {
	err := a.Error()
	return err
}
