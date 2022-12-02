package api

import (
	h "backend/api/handlers"

	"github.com/gin-gonic/gin"
)

func GetRouter() {
	router := gin.Default()
	//router.SetTrustedProxies([]string{"192.168.1.2"})  not safe to trust all proxies (research this)

	user := router.Group("/user")
	{
		user.GET("/test", h.GetString)
		user.GET("/test_endpoint", h.TestEndpoint)
	}

	database := router.Group("/database")
	{
		user.GET("/database", h.GetDatabase)
		database.GET("/viewall", h.ViewDatabase)
	}

	router.Run(":8080")
}
