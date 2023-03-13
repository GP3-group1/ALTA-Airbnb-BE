package helpers

import (
	"alta-airbnb-be/utils/consts"
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LimitOffsetConvert(page, limit int) (int, int) {
	offset := -1
	if limit > 0 {
		offset = (page - 1) * limit
	}
	return limit, offset
}

func ExtractIDParam(c echo.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New(consts.ECHO_InvaildIdParam)
	}
	return id, nil
}

func ExtractPageLimit(c echo.Context) (page int, limit int, err error) {
	pageStr := c.QueryParam("page")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1, errors.New(consts.ECHO_InvaildPageParam)
	}
	limitStr := c.QueryParam("limit")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return -1, -1, errors.New(consts.ECHO_InvaildLimitParam)
	}
	return page, limit, nil
}

func ExtractImage(c echo.Context, key string) (multipart.File, string, error) {
	f, err := c.FormFile(key)
	if err != nil {
		return nil, "", err
	}

	blobFile, err := f.Open()
	if err != nil {
		return nil, "", err
	}
	return blobFile, f.Filename, nil
}
