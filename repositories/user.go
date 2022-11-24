package repositories

import (
	_config "api/mvc/config"
	_models "api/mvc/models"
	"errors"
)

func GetAllUser() ([]_models.User, error) {
	var users []_models.User

	tx := _config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func InsertUser(user _models.User) error {
	tx := _config.DB.Create(&user) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func GetUserbyID(id int) (_models.User, error) {
	var data _models.User
	data.ID = uint(id)

	tx := _config.DB.Take(&data, data.ID)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}

func DeleteUser(id int) error {
	idUser := _models.User{}

	tx := _config.DB.Delete(&idUser, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error delete database by id")
	}
	return nil
}

func UpdateUserbyID(id int, updateReq _models.User) error {
	tx := _config.DB.Where("id = ?", id).Updates(updateReq)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update user failed")
	}
	return nil
}
