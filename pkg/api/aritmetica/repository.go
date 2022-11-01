package aritmetica

import "gorm.io/gorm"

type Repository interface {
	SumNumbers()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) SumNumbers() {

}
