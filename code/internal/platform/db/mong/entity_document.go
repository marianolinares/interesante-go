package mong

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type entityDocument struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Price     int                `bson:"price"`
	Name      string             `bson:"name"`
}
