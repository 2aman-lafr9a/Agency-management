package v1

import (
	models2 "offer/internal/models"
)

type OfferInterface interface {
	Create2(offer *models2.Offer) (models2.Offer, error)
	FindAll2() ([]*models2.Offer, error)
	FindById2(id string) (*models2.Offer, error)
	Update2(offer *models2.Offer) error
	Delete2(offer *models2.Offer) error
	FindByName2(name string) (*models2.Offer, error)
	FindAgencyByID(id string) (*models2.Agency, error)
	FindOfferByID(id string) (*models2.Offer, error)
}

type UseCaseInterface interface {
	Create2(offer *models2.Offer) (models2.Offer, error)
	Update2(offer *models2.Offer) error
	Delete2(offer *models2.Offer) error
	FindAll2() ([]*models2.Offer, error)
	FindById2(id string) (*models2.Offer, error)
	FindByName2(name string) (*models2.Offer, error)
	FindAgencyByID(id string) (*models2.Agency, error)
	FindOfferByID(id string) (*models2.Offer, error)
}
