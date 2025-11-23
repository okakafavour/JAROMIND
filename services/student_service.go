package services

import "github.com/okakafavour/jaromind-backend/models"

type StudentService interface {
	Register(student models.Student) error
	Login(email, password string) (string, error)
}
