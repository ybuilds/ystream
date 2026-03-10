package database

import (
	"context"
	"log"

	"github.com/ybuilds/ystream/backend/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	_ "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Database

func getDb() *mongo.Database {
	return db
}

func LoadDb() error {
	dbUri, err := utils.GetEnvValue("DB_URI")
	if err != nil {
		log.Println("error fetcing database uri, setting to fallback localhost uri")
		dbUri = "mongodb://localhost:27017/ystream"
	}

	clientOptions := options.Client().ApplyURI(dbUri)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Println("error creating database connection")
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("error pinging database")
		return err
	}

	db = client.Database("ystream")

	return nil
}
