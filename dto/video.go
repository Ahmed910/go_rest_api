package dto

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
)

type VideoResponseBody struct {
	ID          uint64 `json:"video_id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func GetVideos(videos []entity.Video) []VideoResponseBody {
	var data []VideoResponseBody
	for _, video := range videos {
		data = append(data, VideoResponseBody{video.ID, video.Title, video.URL, video.Description})
	}
	return data
}
