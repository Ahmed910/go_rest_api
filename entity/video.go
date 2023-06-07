package entity

import (
	"gorm.io/gorm"
	"log"
)

type Video struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title" binding:"min=2,max=10" gorm:"type:varchar(100)"` //validate:"is-cool"
	Description string `json:"description" binding:"max=20" gorm:"type:varchar(200)"`
	URL         string `json:"url" binding:"required,url" gorm:"type:varchar(255);UNIQUE"`
	Author      Author `json:"author" binding:"required" gorm:"foreignkey:AuthorID"`
	AuthorID    uint64 `json:"-"`
}

func Save(authorData Author, videoData Video) (Video, error) {

	NewConnectionToDB()
	db.connection.Create(&authorData)
	videoData.AuthorID = authorData.ID
	err := db.connection.Create(&videoData).Error
	if err != nil {
		return Video{}, err
	}
	db.connection.Limit(1).Find(&videoData.Author, videoData.AuthorID)
	return videoData, nil
}

func FindAll() ([]Video, error) {
	NewConnectionToDB()
	var videos []Video
	err := db.connection.Preload("Author").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func Update(authorData Author, videoData Video, id uint64) (error, Video) {
	NewConnectionToDB()
	var video Video
	if err := db.connection.First(&video, id).Error; err != nil {
		return err, Video{}
	}
	authorData.ID = video.AuthorID
	videoData.ID = id
	videoData.AuthorID = video.AuthorID
	return UpdateVideoService(authorData, videoData)
}

func UpdateVideoService(authorData Author, videoData Video) (error, Video) {
	NewConnectionToDB()
	tx := beginTransaction()
	authorData, err := updateAuthor(tx, authorData)
	if err != nil {
		rollbackTransaction(tx)
		return err, Video{}
	}
	videoData, err = updateVideo(tx, videoData)
	if err != nil {
		rollbackTransaction(tx)
		return err, Video{}
	}
	videoData.Author = authorData
	commitTransaction(tx)
	return nil, videoData

}

func updateAuthor(tx *gorm.DB, authorData Author) (Author, error) {
	err := updateObject(tx, authorData)
	return authorData, err
}
func updateVideo(tx *gorm.DB, videoData Video) (Video, error) {
	err := updateObject(tx, videoData)
	return videoData, err
}

// global
func updateGlobal(data interface{}) error {
	return updateObject(db.connection, data)
}

func updateObject(tx *gorm.DB, obj interface{}) error {
	if err := tx.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

func beginTransaction() *gorm.DB {
	return db.connection.Begin()
}
func rollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
func commitTransaction(tx *gorm.DB) {
	tx.Commit()
}

//end global

func Delete(id uint64) error {
	NewConnectionToDB()
	var video Video
	video.ID = id
	err := db.connection.Delete(&video).Error
	if err != nil {
		log.Fatal("Got an error when delete video. Error: ", err)
		return err
	}
	return nil
}
