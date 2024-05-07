package conditionpg

import (
	modelpg "go-sql/internal/pg/model"
	optionmodel "go-sql/pkg/model/option"

	"github.com/Masterminds/squirrel"
)

type WarehouseInterface interface {
	CommonInterface
}

type warehouseImpl struct {
	Sb    squirrel.SelectBuilder
	Query optionmodel.Options
}

func Warehouse(q optionmodel.Options) WarehouseInterface {
	return warehouseImpl{
		Query: q,
	}
}

func (c warehouseImpl) GetSelectBuilder() squirrel.SelectBuilder {
	return c.Sb
}

func (c warehouseImpl) SelectAll() {
	c.Sb = squirrel.Select("*").From(modelpg.TableNames.Warehouses)
}

func (c warehouseImpl) AssignId() {
	if c.Query.Warehouse.HasID() {
		c.Sb = c.Sb.Where(squirrel.Eq{modelpg.WarehouseTableColumns.ID: c.Query.Warehouse.ID})
	}
}
