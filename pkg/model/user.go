package models

import "time"

type User struct {
	tableName struct{}  `pg:"users"`
	ID        string    `pg:"_id,pk"`
	Name      string    `pg:"name"`
	CreatedAt time.Time `pg:"createdAt"`
	UpdatedAt time.Time `pg:"updatedAt"`
}
