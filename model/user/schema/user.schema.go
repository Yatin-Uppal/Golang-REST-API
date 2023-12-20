package userSchema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId       string             `json:"userId" bson:"userId"`
	FirstName    string             `json:"firstName" bson:"firstName"`
	LastName     string             `json:"lastName" bson:"lastName"`
	Email        string             `json:"email" bson:"email"`
	MobileNumber string             `json:"mobileNumber" bson:"mobileNumber"`
	Password     string             `json:"password" bson:"password,omitempty"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

const USER_SCHEMA_NAME = "users"
