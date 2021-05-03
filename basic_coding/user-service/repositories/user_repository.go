package repositories

import (
	"user-service/models"
	"user-service/drivers/sqllite"
)

type UserRepository struct {
	db *sqllite.SqlLiteDatasource
}

func NewUserRepository(db *sqllite.SqlLiteDatasource) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) RegisterUser(user models.User) *models.User {
	user_data := r.db.InsertUser(user)

	if user_data == nil {
        return nil
    }

	return user_data
}