package conditionpg

import (
	modelpg "go-sql/internal/pg/model"
	optionmodel "go-sql/pkg/model/option"

	"github.com/Masterminds/squirrel"
)

type LocationInterface interface {
	CommonInterface
}

type locationImpl struct {
	Sb    squirrel.SelectBuilder
	Query optionmodel.Options
}

func Location(q optionmodel.Options) LocationInterface {
	return locationImpl{
		Query: q,
	}
}

func (c locationImpl) GetSelectBuilder() squirrel.SelectBuilder {
	return c.Sb
}

func (c locationImpl) SelectAll() {
	c.Sb = squirrel.Select("*").From(modelpg.TableNames.Locations)
}

func (c locationImpl) AssignId() {
	if c.Query.Location.HasID() {
		c.Sb = c.Sb.Where(squirrel.Eq{modelpg.LocationTableColumns.ID: c.Query.Location.ID})
	}
}
