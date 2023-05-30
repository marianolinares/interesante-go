package bootstrap

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func InitMongo(ctx context.Context) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("tasker")
}

func InitPostgres(ctx context.Context) *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URI"))

	if err != nil {
		log.Fatal(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
