package delivery

import (
	"alta-airbnb-be/features/images"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ImageDelivery struct {
	imageService images.ImageService_
}

func New(imageService images.ImageService_) images.ImageDelivery_ {
	return &ImageDelivery{
		imageService: imageService,
	}
}

func (imageDelivery *ImageDelivery) AddImage(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	imageRequest := images.ImageRequest{}
	err := c.Bind(&imageRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.IMAGE_ErrorBindImageData))
	}

	file, fileName, err := helpers.ExtractImage(c, "image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error()))
	}
	imageRequest.Image = file
	imageRequest.ImageName = fileName

	imageEntity := ConvertToEntity(&imageRequest)
	imageEntityOutput, err := imageDelivery.imageService.CreateImage(userId, &imageEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateImageFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	imageResponse := ConvertToResponse(imageEntityOutput)
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.IMAGE_SuccessInsertImageData, imageResponse))
}

func (imageDelivery *ImageDelivery) ModifyImage(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	imageId, err := helpers.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}

	imageRequest := images.ImageRequest{}
	err = c.Bind(&imageRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.IMAGE_ErrorBindImageData))
	}

	file, fileName, err := helpers.ExtractImage(c, "image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error()))
	}
	imageRequest.ID = imageId
	imageRequest.Image = file
	imageRequest.ImageName = fileName

	imageEntity := ConvertToEntity(&imageRequest)
	imageEntityOutput, err := imageDelivery.imageService.ChangeImage(userId, &imageEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateImageFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	imageResponse := ConvertToResponse(imageEntityOutput)
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.IMAGE_SuccessUpdateImageData, imageResponse))
}

func (imageDelivery *ImageDelivery) RemoveImage(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	imageRequest := images.ImageRequest{}
	err := c.Bind(&imageRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.IMAGE_ErrorBindImageData))
	}

	file, fileName, err := helpers.ExtractImage(c, "image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error()))
	}
	imageRequest.Image = file
	imageRequest.ImageName = fileName

	imageEntity := ConvertToEntity(&imageRequest)
	err = imageDelivery.imageService.RemoveImage(userId, &imageEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateImageFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.IMAGE_SuccessDeleteImageData))
}
