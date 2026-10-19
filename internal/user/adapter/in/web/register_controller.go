package web

import (
	"net/http"

	"root_academy_go/internal/user/adapter/in/web/dto"
	"root_academy_go/internal/user/application/port/in"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	registerUser in.RegisterUser
}

func NewRegisterController(registerUser in.RegisterUser) *RegisterController {
	return &RegisterController{registerUser: registerUser}
}

func (c *RegisterController) Handle(ctx *gin.Context) {
	var input dto.Register
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.registerUser.Register(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.UserCreated{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	})
}
