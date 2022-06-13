package service

import (
	"github.com/tris-tux/go-library/backend/repository"
	"github.com/tris-tux/go-library/backend/schema"
)

type Visitor interface {
	FindAll() ([]schema.Visitor, error)
	FindByNoIdentitas(NoIdentitas uint64) (*schema.Visitor, error)
	Create(visitor schema.Visitor) error
	Update(NoIdentitas uint64, visitor schema.Visitor) error
	Delete(NoIdentitas uint64) error
}

type visitor struct {
	repository repository.Repository
}

func NewVisitor(repository repository.Repository) *visitor {
	return &visitor{repository}
}

func (t *visitor) FindAll() ([]schema.Visitor, error) {
	visitors, err := t.repository.FindVisitorAll()
	if err != nil {
		return nil, err
	}
	return visitors, err
}

func (t *visitor) FindByNoIdentitas(NoIdentitas uint64) (*schema.Visitor, error) {
	visitor, err := t.repository.FindVisitorByNoIdentitas(NoIdentitas)

	return visitor, err
}

func (t *visitor) Create(visitor schema.Visitor) error {
	v := schema.Visitor{
		NoIdentitas:     visitor.NoIdentitas,
		Nama:            visitor.Nama,
		TglLahir:        visitor.TglLahir,
		JenisKelamin:    visitor.JenisKelamin,
		Alamat:          visitor.Alamat,
		Handphone:       visitor.Handphone,
		TlpRumah:        visitor.TlpRumah,
		Email:           visitor.Email,
		KdPropinsi:      visitor.KdPropinsi,
		KdKotaKabupaten: visitor.KdKotaKabupaten,
		KdKecamatan:     visitor.KdKecamatan,
		KdKelurahan:     visitor.KdKelurahan,
		Kodepos:         visitor.Kodepos,
		KdJenisId:       visitor.KdJenisId,
		PhotoDiriktp:    visitor.PhotoDiriktp,
	}

	err := t.repository.Create(v)

	return err
}

func (t *visitor) Update(NoIdentitas uint64, visitor schema.Visitor) error {
	v := schema.Visitor{
		NoIdentitas:     visitor.NoIdentitas,
		Nama:            visitor.Nama,
		TglLahir:        visitor.TglLahir,
		JenisKelamin:    visitor.JenisKelamin,
		Alamat:          visitor.Alamat,
		Handphone:       visitor.Handphone,
		TlpRumah:        visitor.TlpRumah,
		Email:           visitor.Email,
		KdPropinsi:      visitor.KdPropinsi,
		KdKotaKabupaten: visitor.KdKotaKabupaten,
		KdKecamatan:     visitor.KdKecamatan,
		KdKelurahan:     visitor.KdKelurahan,
		Kodepos:         visitor.Kodepos,
		KdJenisId:       visitor.KdJenisId,
		PhotoDiriktp:    visitor.PhotoDiriktp,
	}
	err := t.repository.UpdateVisitor(v)

	return err
}

func (t *visitor) Delete(NoIdentitas uint64) error {
	err := t.repository.Delete(NoIdentitas)

	return err
}
