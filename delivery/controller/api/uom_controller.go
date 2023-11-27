package api

import (
	"go_laundry/delivery/middleware"
	"go_laundry/model"
	"go_laundry/model/dto"
	"go_laundry/usecase"
	"go_laundry/util/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UomController struct {
	uomUC usecase.UomUseCase
	rg    *gin.RouterGroup
}

func (u *UomController) createUomHandler(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	uom.Id = common.GenerateID()
	if err := u.uomUC.CreateNew(uom); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, uom)

}

func (u *UomController) updateHandlerUom(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}

	err := u.uomUC.Update(uom)
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Updated Successfully",
	})
}

func (u *UomController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	uoms, paging, err := u.uomUC.Paging(dto.PageRequest{
		Page: page,
		Size: size,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "successfully get uom",
		"data":    uoms,
		"paging":  paging,
	}

	c.JSON(200, response)
}

func (u *UomController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	uom, err := u.uomUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}

	c.JSON(200, uom)
}

func (u *UomController) deleteHandlerUom(c *gin.Context) {
	id := c.Param("id")
	err := u.uomUC.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Uom Deleted Successfully",
	})
}

func (u *UomController) Route() {

	u.rg.POST("/uoms", middleware.AuthMiddleware(), u.createUomHandler)
	u.rg.GET("/uoms", middleware.AuthMiddleware(), u.listHandler)
	u.rg.GET("/uoms/:id", middleware.AuthMiddleware(), u.getByIdHandler)
	u.rg.PUT("/uoms", middleware.AuthMiddleware(), u.updateHandlerUom)
	u.rg.DELETE("/uoms/:id", middleware.AuthMiddleware(), u.deleteHandlerUom)
}

func NewUomController(uomUc usecase.UomUseCase, rg *gin.RouterGroup) *UomController {
	return &UomController{
		uomUC: uomUc,
		rg:    rg,
	}

}
