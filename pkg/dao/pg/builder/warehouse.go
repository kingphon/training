package builderpg

import (
	"context"
	conditionpg "go-sql/pkg/dao/pg/condition"
	optionmodel "go-sql/pkg/model/option"
)

type WarehouseInterface interface {
	CommonInterface
}

type warehouseImpl struct{}

func Warehouse() WarehouseInterface {
	return warehouseImpl{}
}

func (warehouseImpl) FindRawByID(ctx context.Context, id string) (res SQLWithArgs) {
	var (
		q = optionmodel.Options{
			Warehouse: optionmodel.Warehouse{
				CommonStruct: optionmodel.CommonStruct{
					ID: id,
				},
			},
		}
		cond = conditionpg.Warehouse(q)
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
