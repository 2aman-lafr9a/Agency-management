package repository

import (
	"gorm.io/gorm"
	models2 "offer/internal/models"
	_ "offer/pkg/v1"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) FindOfferByID(id string) (*models2.Offer, error) {
	offer := &models2.Offer{}
	err := r.db.Where("id=?", id).First(offer).Error
	return offer, err
}

func (r Repository) FindAgencyByID(id string) (*models2.Agency, error) {
	agency := &models2.Agency{}
	err := r.db.Where("id=?", id).First(agency).Error
	return agency, err
}

func New2(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Create2(offer *models2.Offer) (models2.Offer, error) {
	err := r.db.Create(offer).Error
	return *offer, err
}

func (r Repository) FindAll2() ([]*models2.Offer, error) {
	var offers []*models2.Offer
	err := r.db.Find(&offers).Error
	return offers, err
}

func (r Repository) FindById2(id string) (*models2.Offer, error) {
	offer := &models2.Offer{}
	err := r.db.Where("id=?", id).First(offer).Error
	return offer, err
}

func (r Repository) Update2(offer *models2.Offer) error {
	err := r.db.Save(offer).Error
	return err
}

func (r Repository) Delete2(offer *models2.Offer) error {
	err := r.db.Delete(offer).Error
	return err
}

func (r Repository) FindByName2(name string) (*models2.Offer, error) {
	offer := &models2.Offer{}
	err := r.db.Where("name=?", name).First(offer).Error
	return offer, err
}
