package handler

import (
	"errors"
	"net/http"
	"root_academy_go/internal/domain/repository"
	"root_academy_go/internal/domain/service"
	"root_academy_go/internal/http/dto"

	"github.com/gin-gonic/gin"
)

func RightFomat(c *gin.Context, req *dto.CreateUserInput) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Les données envoyées sont invalides ou incomplètes."})
		return false
	}
	return true
}

func Register(repo repository.UserRepository, c *gin.Context) {
	var req dto.CreateUserInput
	if !RightFomat(c, &req) {
		return
	}

	created, err := service.CreateUser(repo, req)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Cet email est déjà utilisé."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Une erreur est survenue."})
		return
	}

	c.JSON(http.StatusCreated, created)
}
