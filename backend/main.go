package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	controller "github.com/ybuilds/ystream/backend/controllers"
	db "github.com/ybuilds/ystream/backend/database"
	route "github.com/ybuilds/ystream/backend/routes"
	utils "github.com/ybuilds/ystream/backend/utils"
)

func main() {
	fmt.Println("loading .env file...")

	err := utils.LoadEnv()
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}

	fmt.Println("loading database connection")

	err = db.LoadDb()
	if err != nil {
		log.Fatalln("error loading database connection", err)
	}

	fmt.Println("started ystream backend service...")

	router := gin.Default()

	router.GET("/health-check", controller.HealthCheck)

	route.MovieRoute(router)

	err = router.Run("localhost:8000")
	if err != nil {
		log.Fatalln("error starting server", err)
	}
}
