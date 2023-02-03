package api

import (
	h "backend/api/handlers"
	"backend/logic"

	"github.com/gin-gonic/gin"
)

type Client struct {
	//Database *database.Database
	Logic *logic.Logic
}

func (c *Client) CreateRouter() (r *gin.Engine) {
	router := gin.Default()
	//router.SetTrustedProxies([]string{"192.168.1.2"})  not safe to trust all proxies

	user := router.Group("/user")
	{
		user.GET("/test", h.GetString)
		user.GET("/test_endpoint", h.TestEndpoint)
	}

	database := router.Group("/database")
	{
		user.GET("/database", h.GetDatabase)
		database.GET("/viewall", h.ViewDatabase)
		database.GET("/testing_viewall", c.handleViewDatabase)
		database.GET("/testing_logic", c.handleViewDatabase)
	}

	return router

	//router.Run(":8080")
}
