package repositories

import "GoIntoduction/repositories/models"

type UserRepositoryInterface interface {
	GetByEmail(email string) *models.User
}
