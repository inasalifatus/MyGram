package database

import (
	"errors"
	"mygram/config"
	"mygram/models"
)

func InsertPhoto(photo models.Photos) (interface{}, error) {
	// photo := models.Photos{}
	// ctx.Bind(&photo)
	if err := config.DB.Save(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

func GetPhotos() ([]models.Photos, error) {
	var photo []models.Photos
	if err := config.DB.Find(&photo).Error; err != nil {
		return photo, err
	}
	return photo, nil
}

func GetPhotoById(id int) (models.Photos, error) {
	var photo models.Photos

	if rows := config.DB.Find(&photo, id).RowsAffected; rows < 1 {
		err := errors.New("photo not found")
		return photo, err
	}

	return photo, nil
}

func DeletePhotoById(id int) error {
	rows := config.DB.Delete(&models.Photos{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("categories to be deleted is not found")
		return err
	}
	return nil
}
