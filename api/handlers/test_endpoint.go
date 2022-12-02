package handlers

import (
	"github.com/gin-gonic/gin"
)

type test struct {
	ID      int64
	Title   string
	Message string
}

func TestEndpoint(gctx *gin.Context) {
	var temp test
	temp.ID = 10
	temp.Title = "Test endpoint"
	temp.Message = "Hi you have succesfully hit an endpoint :)"

	gctx.JSON(200, temp)
}
