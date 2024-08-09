package dto

import "golang-ddd-demo/domain/entity"

type CatDto struct {
	Id      string  `json:"id"`
	CatName string  `json:"cat_name"`
	Price   float32 `json:"price"`
}

func (g CatDto) MapCatListEntityDto(f []entity.Cat) []CatDto {
	var result []CatDto
	for _, cat := range f {
		result = append(result, CatDto{
			Id:      cat.Id,
			CatName: cat.CatName,
			Price:   cat.Price,
		})
	}

	return result
}
