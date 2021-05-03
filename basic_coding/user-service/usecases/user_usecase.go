package usecase

import (
	"user-service/repositories"
	"user-service/models"
)

type UserUsecase struct {
	rp *repositories.UserRepository
}

func NewUserUsecase(rp *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{rp}
}


func (r *UserUsecase) RegisterUser(user models.User) *models.User {

	status_user := r.rp.RegisterUser(user)

	return status_user
}
