package builderpg

import (
	"context"
	conditionpg "go-sql/pkg/dao/pg/condition"
	optionmodel "go-sql/pkg/model/option"
)

type LocationInterface interface {
	CommonInterface
}

type locationImpl struct{}

func Location() LocationInterface {
	return locationImpl{}
}

func (locationImpl) FindRawByID(ctx context.Context, id string) (res SQLWithArgs) {
	var (
		q = optionmodel.Options{
			Location: optionmodel.Location{
				CommonStruct: optionmodel.CommonStruct{
					ID: id,
				},
			},
		}
		cond = conditionpg.Location(q)
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
