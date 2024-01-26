package v1

import (
	"agency/internal/models"
)

type AgencyInterface interface {
	Create(agency *models.Agency) (models.Agency, error)
	FindAll() ([]*models.Agency, error)
	FindById(id string) (*models.Agency, error)
	Update(agency *models.Agency) error
	Delete(agency *models.Agency) error
	FindByName(name string) (*models.Agency, error)
	FindByEmail(email string) (*models.Agency, error)
	FindOffersByAgencyID(id string) ([]*models.Offer, error)
	FindAgencyByWalletAddress(walletAddress string) (*models.Agency, error)
}

type UseCaseInterface interface {
	Create(agency *models.Agency) (models.Agency, error)
	Update(agency *models.Agency) error
	Delete(agency *models.Agency) error
	FindAll() ([]*models.Agency, error)
	FindById(id string) (*models.Agency, error)
	FindByName(name string) (*models.Agency, error)
	FindByEmail(email string) (*models.Agency, error)
	FindOffersByAgencyID(id string) ([]*models.Offer, error)
	FindAgencyByWalletAddress(walletAddress string) (*models.Agency, error)
}
