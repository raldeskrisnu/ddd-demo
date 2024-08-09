package infrastructure

import (
	"github.com/jinzhu/gorm"
	"golang-ddd-demo/domain/entity"
)

type CatRepositoryImpl struct {
	db *gorm.DB
}

func (r *CatRepositoryImpl) GetCats() ([]entity.Cat, error) {

	var listCats []entity.Cat
	err := r.db.Find(&listCats).Error

	return listCats, err
}
