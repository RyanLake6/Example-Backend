package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (c *Client) handleAPICalls(gctx *gin.Context) {
	users, err := c.Logic.ReturnLogic()

	// Currently a consistent 400 error occurs with any error down stream of this call
	// TODO: set up 400/500 errors depending on the situation
	if err != nil {
		log.Print(err)
		gctx.JSON(400, err)
	} else {
		gctx.JSON(200, users)
	}
}

func (c *Client) handleUserAllAPI(gctx * gin.Context) {
	users, err := c.Logic.ReturnLogic()

	if err != nil {
		log.Print(err)
		gctx.JSON(400, err)
	} else {
		gctx.JSON(200, users)
	}
}
