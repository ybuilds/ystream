package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Movie struct {
	Id          bson.ObjectID `bson:"_id" json:"_id"`
	ImdbId      string        `bson:"imdb_id" json:"imdb_id"`
	Title       string        `bson:"title" json:"title"`
	PosterPath  string        `bson:"poster_path" json:"poster_path"`
	YoutubeId   string        `bson:"youtube_id" json:"youtube_id"`
	Genre       []Genre       `bson:"genre" json:"genre"`
	AdminReview string        `bson:"admin_review" json:"admin_review"`
	Ranking     Ranking       `bson:"ranking" json:"ranking"`
}
