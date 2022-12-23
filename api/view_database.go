package api

import (
	"github.com/gin-gonic/gin"
)

func (c *Client) handleViewDatabase(gctx *gin.Context) {
	users, err := c.Database.GetUser()

	if err != nil {
		// Bad
	}

	gctx.JSON(200, users)

	// gctx.JSON(200, gin.H{
	// 	"message": "Temporarily commented out this api endpoint",
	// })
}
