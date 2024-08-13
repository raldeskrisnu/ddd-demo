package service

import (
	"golang-ddd-demo/domain/dto"
	repository "golang-ddd-demo/infrastructure"
	"log"
	"net/http"
)

type CatServiceImpl struct {
	catRepository repository.CatRepository
}

func (c CatServiceImpl) GetCats() ([]dto.CatDto, dto.Response) {
	var response dto.Response

	cats, err := c.catRepository.GetCats()

	if err != nil {
		response.Status = http.StatusBadRequest
	}

	catsDto := dto.CatDto{}.MapCatListEntityDto(cats)
	response.Status = http.StatusOK
	return catsDto, response
}

func InitCatServiceImpl() *CatServiceImpl {
	dbHelper, err := repository.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &CatServiceImpl{
		catRepository: dbHelper.CatRepository,
	}
}
