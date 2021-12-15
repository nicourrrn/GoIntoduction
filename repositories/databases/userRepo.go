package databases

import "GoIntoduction/repositories/models"

type UserDBRepo struct {
}

func (u UserDBRepo) GetByEmail(email string) *models.User {
	return nil
}
