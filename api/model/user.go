package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            		primitive.ObjectID `json:"userId,omitempty" bson:"_id"`
	UserName          	string             `json:"userName" validate:"required"`
	FirstName          	string             `json:"firstName"`
	Email           	string             `json:"email" validate:"required"`
	Mobile      		string             `json:"mobile" validate:"required"`
	Password          	string             `json:"password" validate:"required"`
	CreationTimestamp 	time.Time          `json:"creationTimestamp,omitempty" bson:"creationTimestamp"`
}
