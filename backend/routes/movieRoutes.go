package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ybuilds/ystream/backend/controllers"
)

func MovieRoute(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())
}
