package helper

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/validation"
)

func BuildAuthorDataFromRequest(videoRequest validation.Video) entity.Author {
	authorData := entity.Author{
		FirstName: videoRequest.Author.FirstName,
		LastName:  videoRequest.Author.LastName,
		Age:       videoRequest.Author.Age,
		Email:     videoRequest.Author.Email,
	}
	return authorData
}

func BuildVideoDataFromRequest(videoRequest validation.Video) entity.Video {
	videoData := entity.Video{
		Title:       videoRequest.Title,
		Description: videoRequest.Description,
		URL:         videoRequest.URL,
	}
	return videoData
}
