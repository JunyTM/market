package repository

import (
	"market/infrastructure"
	"market/model"
)

type UserRepository struct {}

func (r *UserRepository) GetById(id int) (*model.UserRepository, error) {
	db := infrastructure.GetDB()

	if err := db.First(&model.UserRepository).Error; err != nil {
		return nil, err
	}
	return 
}