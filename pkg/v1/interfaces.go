package v1

import "Agency_management/internal/models"

type AgencyInterface interface {
	Create(agency *models.Agency) (models.Agency, error)
	FindAll() ([]*models.Agency, error)
	FindById(id string) (*models.Agency, error)
	Update(agency *models.Agency) error
	Delete(agency *models.Agency) error
	FindByName(name string) (*models.Agency, error)
	FindByEmail(email string) (*models.Agency, error)
}

type OfferInterface interface {
	Create2(offer *models.Offer) (models.Offer, error)
	FindAll2() ([]*models.Offer, error)
	FindById2(id string) (*models.Offer, error)
	Update2(offer *models.Offer) error
	Delete2(offer *models.Offer) error
	FindByName2(name string) (*models.Offer, error)
}

type UseCaseInterface interface {
	Create(agency *models.Agency) (models.Agency, error)
	Update(agency *models.Agency) error
	Delete(agency *models.Agency) error
	FindAll() ([]*models.Agency, error)
	FindById(id string) (*models.Agency, error)
	FindByName(name string) (*models.Agency, error)
	FindByEmail(email string) (*models.Agency, error)
	Create2(offer *models.Offer) (models.Offer, error)
	Update2(offer *models.Offer) error
	Delete2(offer *models.Offer) error
	FindAll2() ([]*models.Offer, error)
	FindById2(id string) (*models.Offer, error)
	FindByName2(name string) (*models.Offer, error)
}
