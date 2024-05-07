package conditionpg

import (
	modelpg "go-sql/internal/pg/model"
	optionmodel "go-sql/pkg/model/option"

	"github.com/Masterminds/squirrel"
)

type UserInterface interface {
	CommonInterface
}

type userImpl struct {
	Sb    squirrel.SelectBuilder
	Query optionmodel.Options
}

func User(q optionmodel.Options) UserInterface {
	return userImpl{
		Query: q,
	}
}

func (c userImpl) GetSelectBuilder() squirrel.SelectBuilder {
	return c.Sb
}

func (c userImpl) SelectAll() {
	c.Sb = squirrel.Select("*").From(modelpg.TableNames.Users)
}

func (c userImpl) AssignId() {
	if c.Query.User.HasID() {
		c.Sb = c.Sb.Where(squirrel.Eq{modelpg.UserTableColumns.ID: c.Query.User.ID})
	}
}
