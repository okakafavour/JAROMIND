package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okakafavour/jaromind-backend/models"
	servicesimpl "github.com/okakafavour/jaromind-backend/services_impl"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterStudent(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err := servicesimpl.NewStudentService().Register(student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful! Please check your email to verify your account.",
	})
}

type LoginRequest struct {
	Email    string `bson:"email" binding:"required,email"`
	Password string `bson:"password" binding:"required"`
}

func LoginStudent(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := servicesimpl.NewStudentService().Login(request.Email, request.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successfully",
		"token":   token,
	})
}
