package api

import (
	"go_laundry/model/dto"
	"go_laundry/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUC usecase.UserUseCase
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var dto dto.AuthRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	authResponse, err := a.authUC.Login(dto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully Login",
		"date":    authResponse,
	}
	c.JSON(200, response)
}

func (a *AuthController) registerHandler(c *gin.Context) {
	var dto dto.AuthRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := a.userUC.Register(dto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully Register",
	}
	c.JSON(200, response)
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/register", a.registerHandler)
}

func NewAuthController(userUC usecase.UserUseCase, authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		userUC: userUC,
		authUC: authUC,
		rg:     rg,
	}
}
