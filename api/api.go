package api

import (
	"backend/logic"

	"github.com/gin-gonic/gin"
)

type Client struct {
	Logic *logic.Logic
}

func (c *Client) CreateRouter() (r *gin.Engine) {
	router := gin.Default()
	//router.SetTrustedProxies([]string{"192.168.1.2"})  not safe to trust all proxies

	database := router.Group("/database")
	{
		database.GET("/testing_logic", c.handleAPICalls)
		database.GET("/test", c.handleAPICalls)
	}

	user := router.Group("/user") 
	{
		user.GET("/all", c.handleUserAllAPI)
	}

	return router
}
