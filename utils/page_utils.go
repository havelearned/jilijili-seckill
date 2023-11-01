package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PageDto struct {
	Current int         `form:"current" json:"current"`
	Size    int         `form:"size" json:"size"`
	Records interface{} `form:"records" json:"records"`
	Total   int64       `form:"total" json:"total"`
}

func GetPageAndSize(requestForm *gin.Context) (int, int) {
	page := requestForm.Query("page")
	size := requestForm.Query("size")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		sizeInt = 10
	}

	return pageInt, sizeInt
}
