package api

import (
	"go_laundry/model"
	"go_laundry/model/dto"
	"go_laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeController struct {
	employeeUC usecase.EmployeeUseCase
	rg         *gin.RouterGroup
}

func (e *EmployeeController) createHandlerEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}

	employee.Id = uuid.NewString()
	if err := e.employeeUC.CreateNew(employee); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (e *EmployeeController) updateHandlerEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
		return
	}

	err := e.employeeUC.Update(employee)
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

func (e *EmployeeController) listHandlerEmployee(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	employees, paging, err := e.employeeUC.Paging(dto.PageRequest{
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
		"message": "successfully get Employee",
		"data":    employees,
		"paging":  paging,
	}

	c.JSON(200, response)
}

func (e *EmployeeController) getByIdHandlerEmployee(c *gin.Context) {
	id := c.Param("id")
	employee, err := e.employeeUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}

	c.JSON(200, employee)
}

func (e *EmployeeController) deleteHandlerEmployee(c *gin.Context) {
	id := c.Param("id")
	err := e.employeeUC.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Employee Deleted Successfully",
	})
}

func (e *EmployeeController) Route() {

	e.rg.POST("/employees", e.createHandlerEmployee)
	e.rg.GET("/employees", e.listHandlerEmployee)
	e.rg.GET("/employees/:id", e.getByIdHandlerEmployee)
	e.rg.PUT("/employees", e.updateHandlerEmployee)
	e.rg.DELETE("/employees/:id", e.deleteHandlerEmployee)
}

func NewEmployeeController(employeeUC usecase.EmployeeUseCase, rg *gin.RouterGroup) *EmployeeController {
	return &EmployeeController{
		employeeUC: employeeUC,
		rg:         rg,
	}

}
