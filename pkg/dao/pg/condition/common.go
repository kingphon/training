package conditionpg

import (
	"github.com/Masterminds/squirrel"
)

type CommonInterface interface {
	GetSelectBuilder() squirrel.SelectBuilder
	SelectAll()

	AssignId()
}
