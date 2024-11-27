package routes

import (
	"event-booking-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(router *gin.Context) {
	var user models.User

	if err := router.ShouldBindJSON(&user); err != nil {
		router.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := user.Save()
	if err != nil {
		router.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	router.JSON(http.StatusCreated, createdUser)
}

func login(router *gin.Context) {
	var user models.User

	if err := router.ShouldBindJSON(&user); err != nil {
		router.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.ValidateCredentials()
	if err != nil {
		router.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	router.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
