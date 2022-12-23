package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (c *Client) handleViewDatabase(gctx *gin.Context) {
	users, err := c.Database.GetUser()

	if err != nil {
		log.Fatal("Hit error: ", err)
		// Bad
	}

	gctx.JSON(200, users)

	// gctx.JSON(200, gin.H{
	// 	"message": "Temporarily commented out this api endpoint",
	// })
}
