package postgres

import "time"

const sqlEntityTable = "entity"

type entityDocument struct {
	ID        string    `db:"_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Price     int       `db:"price"`
	Name      string    `db:"named"`
}
