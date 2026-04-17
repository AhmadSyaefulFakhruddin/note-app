package tags

import (
	"net/http"
	"note-app-api/internal/features/apperr"
	"note-app-api/internal/features/response"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) GetTags(c *gin.Context) {
	tagsRes, err := h.service.GetTags(c.Request.Context())

	if err != nil {
		errCode, errRes := response.Error(err)

		c.JSON(errCode, errRes)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(tagsRes, "Successfully to get tags"))
}

func (h *handler) CreateMultipleTags(c *gin.Context) {
	var createMultipleTagsDto CreateMultipleTagsDto

	if err := c.ShouldBindBodyWithJSON(&createMultipleTagsDto); err != nil {
		errCode, errRes := response.Error(apperr.NewInternal(err))

		c.JSON(errCode, errRes)
		return
	}

	tagDtos, err := h.service.CreateMultipleTags(c, createMultipleTagsDto.TagNames)

	if err != nil {
		errCode, errRes := response.Error(err)

		c.JSON(errCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(tagDtos, "Successfully to create Create multiple tags"))
}
