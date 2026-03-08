package models

type Genre struct {
	GenreId   int    `bson:"genre_id" json:"genre_id"`
	GenreName string `bson:"genre_name" json:"genre_name"`
}
