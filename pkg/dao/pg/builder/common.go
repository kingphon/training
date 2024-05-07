package builderpg

import (
	"context"
)

type SQLWithArgs struct {
	SQL  string
	Args []interface{}
	Err  error
}

type CommonInterface interface {
	FindRawByID(ctx context.Context, id string) SQLWithArgs
	// FindRawsByCondition(ctx context.Context, q optionmodel.Options) SQLWithArgs
}
