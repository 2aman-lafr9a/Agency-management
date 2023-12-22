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
	var agencies []*models.Agency
	err := r.db.Find(&agencies).Error
	return agencies, err
}

func (r Repository) FindById(id string) (*models.Agency, error) {
	agency := &models.Agency{}
	err := r.db.Where("id=?", id).First(agency).Error
	return agency, err
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
	agency := &models.Agency{}
	err := r.db.Where("name=?", name).First(agency).Error
	return agency, err
}

func (r Repository) FindByEmail(email string) (*models.Agency, error) {
	err := r.db.Where("email=?", email).First(&models.Agency{}).Error
	return nil, err
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func New2(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Create2(offer *models.Offer) (models.Offer, error) {
	err := r.db.Create(offer).Error
	return *offer, err
}

func (r Repository) FindAll2() ([]*models.Offer, error) {
	var offers []*models.Offer
	err := r.db.Find(&offers).Error
	return offers, err
}

func (r Repository) FindById2(id string) (*models.Offer, error) {
	offer := &models.Offer{}
	err := r.db.Where("id=?", id).First(offer).Error
	return offer, err
}

func (r Repository) Update2(offer *models.Offer) error {
	err := r.db.Save(offer).Error
	return err
}

func (r Repository) Delete2(offer *models.Offer) error {
	err := r.db.Delete(offer).Error
	return err
}

func (r Repository) FindByName2(name string) (*models.Offer, error) {
	offer := &models.Offer{}
	err := r.db.Where("name=?", name).First(offer).Error
	return offer, err
}
