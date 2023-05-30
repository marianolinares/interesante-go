package mong

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"marian.com/interesante-go/code/internal"
	"time"
)

type repo struct {
	ctx        context.Context
	collection *mongo.Collection
}

func (r *repo) GetEntities() ([]internal.Entity, error) {
	var anError error = nil
	//var anError error = errors.New("error")

	return []internal.Entity{internal.NewEntity(10, "Pant")}, anError
}

func (r *repo) SaveEntity(e internal.Entity) {
	doc := &entityDocument{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Price:     e.Price,
		Name:      e.Name,
	}

	r.collection.InsertOne(r.ctx, doc)
}

func NewEntityRepository(ctx context.Context, db *mongo.Database) internal.EntityRepo {
	return &repo{
		ctx:        ctx,
		collection: db.Collection("Entity"),
	}
}
