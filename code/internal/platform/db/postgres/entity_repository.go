package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"marian.com/interesante-go/code/internal"
	"time"
)

type repo struct {
	ctx context.Context
	db  *sql.DB
}

func NewEntityRepository(ctx context.Context, db *sql.DB) internal.EntityRepo {
	return &repo{ctx, db}
}

func (r *repo) GetEntities() ([]internal.Entity, error) {
	var anError error = nil
	//var anError error = errors.New("error")

	return []internal.Entity{internal.NewEntity(10, "Pant")}, anError
}

func (r *repo) SaveEntity(e internal.Entity) {
	fmt.Println("Save entity", e)

	entitySqlStruct := sqlbuilder.NewStruct(new(entityDocument))
	query, args := entitySqlStruct.InsertInto(sqlEntityTable, entityDocument{
		ID:        primitive.NewObjectID().Hex(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Price:     e.Price,
		Name:      e.Name,
	}).Build()

	fmt.Println(query)
	fmt.Println(args)

	//_, err := r.db.ExecContext(r.ctx, query, args...)

	sql := "INSERT INTO entity(_id, created_at, updated_at, price, named) VALUES ($1, $2, $3, $4, $5);"
	_, err := r.db.ExecContext(r.ctx, sql, args...)

	if err != nil {
		//return fmt.Errorf("Error trying to persist entity in database: %v", err)
		fmt.Println("Error", err)
	}

	//return nil
	//panic("implement me")
}
