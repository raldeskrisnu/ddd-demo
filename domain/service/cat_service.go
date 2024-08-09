package service

import "golang-ddd-demo/domain/dto"

type CatService interface {
	GetCats() ([]dto.CatDto, dto.Response)
}
