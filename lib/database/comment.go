package database

import (
	"errors"
	"mygram/config"
	"mygram/models"
)

func InsertComment(comment models.Comment) (interface{}, error) {
	// photo := models.Photos{}
	// ctx.Bind(&photo)
	if err := config.DB.Save(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetComment() ([]models.Comment, error) {
	var comment []models.Comment
	if err := config.DB.Find(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

func GetCommentById(id int) (models.Comment, error) {
	var comment models.Comment

	if rows := config.DB.Find(&comment, id).RowsAffected; rows < 1 {
		err := errors.New("comment not found")
		return comment, err
	}

	return comment, nil
}

func DeleteCommentById(id int) error {
	rows := config.DB.Delete(&models.Comment{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("comment to be deleted is not found")
		return err
	}
	return nil
}
