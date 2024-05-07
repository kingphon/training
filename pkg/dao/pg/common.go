package pgdao

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CommonInterface[ModelType any] interface {
	FindRawByID(ctx context.Context, db boil.ContextExecutor, id string) (doc ModelType, err error)
	InsertOne(ctx context.Context, db boil.ContextExecutor, doc ModelType) (err error)
}
