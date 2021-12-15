package filesystem

import "GoIntoduction/repositories/models"

type UserFileRepo struct {
}

func (userRepo UserFileRepo) GetByEmail(email string) *models.User {
	return nil
}
