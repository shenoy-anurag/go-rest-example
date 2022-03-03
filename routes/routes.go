package routes

import (
	"github.com/shenoy-anurag/go-rest-example/users"

	"github.com/shenoy-anurag/go-rest-example/posts"

	"github.com/shenoy-anurag/go-rest-example/communities"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// create gin router server
	router := gin.Default()

	router.Use(cors.Default())
	// ping route as a health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// add various routes to the gin server
	users.Routes(router)
	communities.Routes(router)
	posts.Routes(router)

	return router
}
