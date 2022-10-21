package database

import (
	"errors"
	"mygram/config"
	"mygram/models"
)

func CreateSosmed(sosmed models.Sosmed) (interface{}, error) {
	if err := config.DB.Save(&sosmed).Error; err != nil {
		return nil, err
	}
	return sosmed, nil
}

func GetSosmed() ([]models.Sosmed, error) {
	var sosmed []models.Sosmed
	if err := config.DB.Find(&sosmed).Error; err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func GetSosmedById(id int) (models.Sosmed, error) {
	var sosmed models.Sosmed

	if rows := config.DB.Find(&sosmed, id).RowsAffected; rows < 1 {
		err := errors.New("sosmed not found")
		return sosmed, err
	}

	return sosmed, nil
}

func DeleteSosmedById(id int) error {
	rows := config.DB.Delete(&models.Sosmed{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("sosmed to be deleted is not found")
		return err
	}
	return nil
}
