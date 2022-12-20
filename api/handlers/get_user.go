package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetString(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"message": "Hello world this is the second iteration",
	})
}
