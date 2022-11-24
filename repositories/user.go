package repositories

import (
	"api/mvc/config"
	"api/mvc/models"
	"errors"
)

func GetAllUser() ([]models.User, error) {
	var users []models.User

	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func InsertUser(user models.User) error {
	tx := config.DB.Create(&user) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}
