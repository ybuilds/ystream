package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ybuilds/ystream/backend/controllers"
	"github.com/ybuilds/ystream/backend/utils"
)

func main() {
	fmt.Println("loading .env file...")

	err := utils.LoadEnv()
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}

	fmt.Println("started ystream backend service...")

	router := gin.Default()

	router.GET("/health-check", controllers.HealthCheck)

	err = router.Run("localhost:8000")
	if err != nil {
		log.Fatalln("error starting server", err)
	}
}
