package conditionpg

import (
	modelpg "go-sql/internal/pg/model"
	optionmodel "go-sql/pkg/model/option"

	"github.com/Masterminds/squirrel"
)

type ShopInterface interface {
	CommonInterface
}

type shopImpl struct {
	Sb    squirrel.SelectBuilder
	Query optionmodel.Options
}

func Shop(q optionmodel.Options) ShopInterface {
	return shopImpl{
		Query: q,
	}
}

func (c shopImpl) GetSelectBuilder() squirrel.SelectBuilder {
	return c.Sb
}

func (c shopImpl) SelectAll() {
	c.Sb = squirrel.Select("*").From(modelpg.TableNames.Shops)
}

func (c shopImpl) AssignId() {
	if c.Query.Shop.HasID() {
		c.Sb = c.Sb.Where(squirrel.Eq{modelpg.ShopTableColumns.ID: c.Query.Shop.ID})
	}
}
