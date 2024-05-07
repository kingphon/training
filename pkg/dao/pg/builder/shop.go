package builderpg

import (
	"context"
	conditionpg "go-sql/pkg/dao/pg/condition"
	optionmodel "go-sql/pkg/model/option"
)

type ShopInterface interface {
	CommonInterface
}

type shopImpl struct{}

func Shop() ShopInterface {
	return shopImpl{}
}

func (shopImpl) FindRawByID(ctx context.Context, id string) (res SQLWithArgs) {
	var (
		q = optionmodel.Options{
			Shop: optionmodel.Shop{
				CommonStruct: optionmodel.CommonStruct{
					ID: id,
				},
			},
		}
		cond = conditionpg.Shop(q)
	)

	// Select
	cond.SelectAll()

	// Assign
	cond.AssignId()

	sb := cond.GetSelectBuilder()

	sql, args, err := sb.ToSql()

	return SQLWithArgs{
		SQL:  sql,
		Args: args,
		Err:  err,
	}
}
