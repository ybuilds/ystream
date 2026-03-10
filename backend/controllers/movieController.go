package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ybuilds/ystream/backend/database"
	"github.com/ybuilds/ystream/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(OK, gin.H{"message": "movies"})

		c, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		var movies []models.Movie

		movieCollection, err := database.GetCollection("movies")
		if err != nil {
			ctx.JSON(ISE, gin.H{"error": "error finding collection: " + err.Error()})
			return
		}

		cursor, err := movieCollection.Find(c, bson.M{})
		if err != nil {
			ctx.JSON(ISE, gin.H{"error": "error fetching movie documents: " + err.Error()})
			return
		}

		err = cursor.All(c, &movies)
		if err != nil {
			ctx.JSON(ISE, gin.H{"error": "failed to parse cursor data to variable: " + err.Error()})
			return
		}

		ctx.JSON(OK, gin.H{"movies": movies})
	}
}
