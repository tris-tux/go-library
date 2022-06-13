package repository

import (
	"github.com/tris-tux/go-library/backend/schema"
	"gorm.io/gorm"
)

type Repository interface {
	FindVisitorAll() ([]schema.Visitor, error)
	FindVisitorByNoIdentitas(NoIdentitas uint64) (*schema.Visitor, error)
	Create(visitor schema.Visitor) error
	UpdateVisitor(visitor schema.Visitor) error
	Delete(NoIdentitas uint64) error
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindVisitorAll() ([]schema.Visitor, error) {
	var visitors []schema.Visitor
	err := r.db.Find(&visitors).Error

	if err != nil {
		return nil, schema.ErrorWrap(503, err.Error())
	}

	return visitors, nil
}

func (r *repository) FindVisitorByNoIdentitas(NoIdentitas uint64) (*schema.Visitor, error) {
	var visitor schema.Visitor
	err := r.db.Find(&visitor, NoIdentitas).Error
	if visitor.NoIdentitas == 0 {
		return nil, schema.ErrorWrap(404, "Data Not Found")
	}
	if err != nil {
		return nil, schema.ErrorWrap(503, err.Error())
	}

	return &visitor, nil
}

func (r *repository) Create(visitor schema.Visitor) error {
	err := r.db.Create(&visitor).Error

	return err
}

func (r *repository) UpdateVisitor(visitor schema.Visitor) error {
	err := r.db.Save(&visitor).Error

	return err
}

func (r *repository) Delete(NoIdentitas uint64) error {
	var visitor schema.Visitor
	err := r.db.Find(&visitor, NoIdentitas).Error

	return err
}
