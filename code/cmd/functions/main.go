package main

import (
	"context"
	"log"
	"marian.com/interesante-go/code/cmd/bootstrap"
	"marian.com/interesante-go/code/cmd/functions/async"
	"marian.com/interesante-go/code/cmd/functions/events"
	"marian.com/interesante-go/code/cmd/functions/healthcheck"
	"marian.com/interesante-go/code/internal/platform/db/mong"
	"marian.com/interesante-go/code/internal/platform/db/postgres"
	"marian.com/interesante-go/code/internal/registerStock"
	"net/http"
	"os"
)

func main() {
	listenAddr := ":8082"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	ctx := context.TODO()

	mongodb := bootstrap.InitMongo(ctx)
	postgresql := bootstrap.InitPostgres(ctx)
	defer postgresql.Close()

	repoMongo := mong.NewEntityRepository(ctx, mongodb)
	repoSql := postgres.NewEntityRepository(ctx, postgresql)

	useCase := registerStock.NewUseCase(repoMongo)

	http.HandleFunc("/api/healthcheck", healthcheck.CreateHandler())
	http.HandleFunc("/api/v1/events", events.CreateHandler(repoSql))
	http.HandleFunc("/async", async.CreateHandler(useCase))

	log.Printf("About to listen on %s. Go to http://localhost%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
