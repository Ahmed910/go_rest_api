package service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/helper"
	"gilab.com/pragmaticreviews/golang-gin-poc/validation"
)

func Save(videoRequest validation.Video) (entity.Video, error) {
	var authorData entity.Author = helper.BuildAuthorDataFromRequest(videoRequest)
	var videoData entity.Video = helper.BuildVideoDataFromRequest(videoRequest)
	data, err := entity.Save(authorData, videoData)
	return data, err
}

func FindAll() ([]entity.Video, error) {
	data, err := entity.FindAll()
	return data, err
}

func Update(videoRequest validation.Video, id uint64) (error, entity.Video) {
	var authorData entity.Author = helper.BuildAuthorDataFromRequest(videoRequest)
	var videoData entity.Video = helper.BuildVideoDataFromRequest(videoRequest)
	err, data := entity.Update(authorData, videoData, id)
	return err, data
}
func Delete(id uint64) error {
	err := entity.Delete(id)
	return err
}
