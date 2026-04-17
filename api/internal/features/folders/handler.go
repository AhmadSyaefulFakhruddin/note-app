package folders

import (
	"net/http"
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

func (h *handler) GetFolders(c *gin.Context) {
	folderResponses, err := h.service.GetFolders(c.Request.Context())

	if err != nil {
		errorCode, errorRes := response.Error(err)

		c.JSON(errorCode, errorRes)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(folderResponses, "Successfully to get folders"))
}

func (h *handler) CreateFolder(c *gin.Context) {
	var folderName CreateFolderRequest

	if err := c.ShouldBindBodyWithJSON(&folderName); err != nil {
		errorCode, errorRes := response.Error(err)

		c.JSON(errorCode, errorRes)
		return
	}

	folderResponse, err := h.service.CreateFolder(c, folderName.FolderName)

	if err != nil {
		errorCode, errorRes := response.Error(err)

		c.JSON(errorCode, errorRes)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(folderResponse, "Successfully to create folder"))
}
