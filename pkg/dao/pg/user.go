package pgdao

import (
	modelpg "go-sql/internal/pg/model"
)

type UserInterface interface {
	CommonInterface[*modelpg.User]
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}
