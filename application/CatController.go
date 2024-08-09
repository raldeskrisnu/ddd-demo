package application

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-demo/domain/dto"
	"golang-ddd-demo/domain/service"
	"net/http"
)

type CatController struct {
	catService service.CatService
}

func InitCatController(router *gin.Engine) {
	catController := CatController{
		catService: service.InitCatServiceImpl(),
	}
	router.GET("/cats", catController.GetCatsHandler)
}

func (r *CatController) GetCatsHandler(c *gin.Context) {

	var response dto.Response

	foods, response := r.catService.GetCats()

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, foods)
}
