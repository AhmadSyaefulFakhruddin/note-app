package folders

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *handler) {
	folderRoute := r.Group("/folders")

	{
		folderRoute.GET("", h.GetFolders)

		folderRoute.POST("", h.CreateFolder)
	}
}
