package UseCase

import (
	"agency/internal/models"
	interfaces "agency/pkg/v1"
	_ "errors"
	_ "gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.AgencyInterface
}

func (u UseCase) FindAgencyByWalletAddress(walletAddress string) (*models.Agency, error) {
	agency, err := u.repo.FindAgencyByWalletAddress(walletAddress)
	return agency, err
}

func (u UseCase) FindOffersByAgencyID(id string) ([]*models.Offer, error) {
	agencies, err := u.repo.FindOffersByAgencyID(id)
	return agencies, err
}

func (u UseCase) FindAll() ([]*models.Agency, error) {
	agencies, err := u.repo.FindAll()
	return agencies, err
}

func (u UseCase) FindByName(name string) (*models.Agency, error) {
	agency, err := u.repo.FindByName(name)
	return agency, err
}

func (u UseCase) FindByEmail(email string) (*models.Agency, error) {
	err, _ := u.repo.FindByEmail(email)
	return nil, err
}

func (u UseCase) FindById(id string) (*models.Agency, error) {
	agency, err := u.repo.FindById(id)
	return agency, err
}

func (u UseCase) Create(agency *models.Agency) (models.Agency, error) {
	agenc, err := u.repo.Create(agency)
	return agenc, err
}

func (u UseCase) Update(agency *models.Agency) error {
	err := u.repo.Update(agency)
	return err
}

func (u UseCase) Delete(agency *models.Agency) error {
	err := u.repo.Delete(agency)
	return err
}

func New(repo interfaces.AgencyInterface) interfaces.UseCaseInterface {
	return &UseCase{repo: repo}
}
