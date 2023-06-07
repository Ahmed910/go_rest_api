package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/response"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"gilab.com/pragmaticreviews/golang-gin-poc/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

var validate *validator.Validate

func Save(ctx *gin.Context) {
	var videoRequest validation.Video
	err := ctx.ShouldBindJSON(&videoRequest)
	if err != nil {
		errors, firstError := response.GetErrors(err)
		ctx.JSON(http.StatusUnprocessableEntity, response.BuildValidationErrorResponse(false, firstError, response.Null(), errors))
		return
	}
	data, err := service.Save(videoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BuildResponse(false, err.Error(), response.Null()))
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponse(true, "Video Added Successfully", data))
}

func FindAll(ctx *gin.Context) {
	data, err := service.FindAll()
	if err != nil {
		log.Fatal("Happened Error when find all video data. Error: ", err)
		ctx.JSON(http.StatusInternalServerError, response.BuildResponse(false, err.Error(), response.Null()))
	} else {
		ctx.JSON(http.StatusOK, response.BuildResponse(true, "", dto.GetVideos(data)))
	}

}

func Update(ctx *gin.Context) {
	var videoRequest validation.Video
	err := ctx.ShouldBindJSON(&videoRequest)
	if err != nil {
		errors, firstError := response.GetErrors(err)
		ctx.JSON(http.StatusUnprocessableEntity, response.BuildValidationErrorResponse(false, firstError, response.Null(), errors))
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.BuildResponse(false, err.Error(), response.Null()))
		return
	}
	err, data := service.Update(videoRequest, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BuildResponse(false, err.Error(), response.Null()))
		return
	}
	ctx.JSON(http.StatusOK, response.BuildResponse(true, "Video Updated Successfully", data))
}
func Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.BuildResponse(false, err.Error(), response.Null()))
		return
	}
	err = service.Delete(id)
	if err != nil {
		log.Fatal("Happened Error when update video data. Error: ", err)
		ctx.JSON(http.StatusInternalServerError, response.BuildResponse(false, err.Error(), response.Null()))
		return
	}
	ctx.JSON(http.StatusOK, response.BuildResponse(true, "Video Deleted Successfully", response.Null()))
}
