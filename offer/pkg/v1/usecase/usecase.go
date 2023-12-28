package UseCase

import (
	_ "errors"
	_ "gorm.io/gorm"
	models2 "offer/internal/models"
	interfaces "offer/pkg/v1"
)

type UseCase struct {
	repo2 interfaces.OfferInterface
}

func New2(repo interfaces.OfferInterface) interfaces.UseCaseInterface {
	return &UseCase{repo2: repo}
}

func (u UseCase) FindAll2() ([]*models2.Offer, error) {
	offers, err := u.repo2.FindAll2()
	return offers, err
}

func (u UseCase) FindByName2(name string) (*models2.Offer, error) {
	offer, err := u.repo2.FindByName2(name)
	return offer, err
}

func (u UseCase) FindById2(id string) (*models2.Offer, error) {
	offer, err := u.repo2.FindById2(id)
	return offer, err
}

func (u UseCase) Create2(offer *models2.Offer) (models2.Offer, error) {
	err, _ := u.repo2.Create2(offer)
	return *offer, err
}

func (u UseCase) Update2(offer *models2.Offer) error {
	err := u.repo2.Update2(offer)
	return err
}

func (u UseCase) Delete2(offer *models2.Offer) error {
	err := u.repo2.Delete2(offer)
	return err
}
