package UseCase

import (
	"Agency_management/internal/models"
	interfaces "Agency_management/pkg/v1"
	_ "errors"
	_ "gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.AgencyInterface
}

func (u UseCase) GetById(id string) (*models.Agency, error) {
	_, err := u.repo.FindById(id)
	return nil, err
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
	_, err := u.repo.FindById(id)
	return nil, err
}

func (u UseCase) Create(agency *models.Agency) (models.Agency, error) {
	err, _ := u.repo.Create(agency)
	return *agency, err
}

func (u UseCase) GetAll() ([]*models.Agency, error) {
	_, err := u.repo.FindAll()
	return nil, err
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
