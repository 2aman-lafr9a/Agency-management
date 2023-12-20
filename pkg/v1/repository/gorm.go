package repository

import (
	"Agency_management/internal/models"
	_ "Agency_management/pkg/v1"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) Create(agency *models.Agency) (models.Agency, error) {
	err := r.db.Create(agency).Error
	return *agency, err
}

func (r Repository) FindAll() ([]*models.Agency, error) {
	err := r.db.Find(&models.Agency{}).Error
	return nil, err
}

func (r Repository) FindById(id string) (*models.Agency, error) {
	err := r.db.First(&models.Agency{}, id).Error
	return nil, err
}

func (r Repository) Update(agency *models.Agency) error {
	err := r.db.Save(agency).Error
	return err
}

func (r Repository) Delete(agency *models.Agency) error {
	err := r.db.Delete(agency).Error
	return err
}

func (r Repository) FindByName(name string) (*models.Agency, error) {
	err := r.db.Where("name=?", name).First(&models.Agency{}).Error
	return nil, err
}

func (r Repository) FindByEmail(email string) (*models.Agency, error) {
	err := r.db.Where("email=?", email).First(&models.Agency{}).Error
	return nil, err
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
