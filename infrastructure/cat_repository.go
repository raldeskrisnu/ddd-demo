package infrastructure

import "golang-ddd-demo/domain/entity"

type CatRepository interface {
	GetCats() ([]entity.Cat, error)
}
