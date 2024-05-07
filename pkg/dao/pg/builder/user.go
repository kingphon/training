package builderpg

import (
	"context"
	conditionpg "go-sql/pkg/dao/pg/condition"
	optionmodel "go-sql/pkg/model/option"
)

type UserInterface interface {
	CommonInterface
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

func (userImpl) FindRawByID(ctx context.Context, id string) (res SQLWithArgs) {
	var (
		q = optionmodel.Options{
			User: optionmodel.User{
				CommonStruct: optionmodel.CommonStruct{
					ID: id,
				},
			},
		}
		cond = conditionpg.User(q)
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

func (userImpl) FindRawsByCondition(ctx context.Context, q optionmodel.Options) (res SQLWithArgs) {
	var (
		cond = conditionpg.User(q)
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
