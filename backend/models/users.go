package models

import "time"

type User struct {
	UserId       int       `bson:"user_id" json:"user_id"`
	FirstName    string    `bson:"first_name" json:"first_name"`
	LastName     string    `bson:"last_name" json:"last_name"`
	Email        string    `bson:"email" json:"email"`
	Password     string    `bson:"password" json:"password"`
	Role         string    `bson:"role" json:"role"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
	Token        string    `bson:"token" json:"imdtokenb_id"`
	RefreshToken string    `bson:"refresh_token" json:"refresh_token"`
	Genre        []Genre   `bson:"favourite_genres" json:"favourite_genres"`
}
