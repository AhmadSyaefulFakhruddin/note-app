package tags

import "github.com/gin-gonic/gin"

func RegisterRoute(c *gin.Engine, h *handler) {

	tagsGroupRoute := c.Group("/tags")

	{
		tagsGroupRoute.GET("", h.GetTags)

		tagsGroupRoute.POST("", h.CreateMultipleTags)
	}

}
