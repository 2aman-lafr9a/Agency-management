package UseCase

import (
	"Agency_management/internal/models"
	interfaces "Agency_management/pkg/v1"
	_ "errors"
	_ "gorm.io/gorm"
)

type UseCase struct {
	repo  interfaces.AgencyInterface
	repo2 interfaces.OfferInterface
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
	err, _ := u.repo.Create(agency)
	return *agency, err
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

func New2(repo interfaces.OfferInterface) interfaces.UseCaseInterface {
	return &UseCase{repo2: repo}
}

func (u UseCase) FindAll2() ([]*models.Offer, error) {
	offers, err := u.repo2.FindAll2()
	return offers, err
}

func (u UseCase) FindByName2(name string) (*models.Offer, error) {
	offer, err := u.repo2.FindByName2(name)
	return offer, err
}

func (u UseCase) FindById2(id string) (*models.Offer, error) {
	offer, err := u.repo2.FindById2(id)
	return offer, err
}

func (u UseCase) Create2(offer *models.Offer) (models.Offer, error) {
	err, _ := u.repo2.Create2(offer)
	return *offer, err
}

func (u UseCase) Update2(offer *models.Offer) error {
	err := u.repo2.Update2(offer)
	return err
}

func (u UseCase) Delete2(offer *models.Offer) error {
	err := u.repo2.Delete2(offer)
	return err
}
