package pgdao

import (
	modelpg "go-sql/internal/pg/model"
)

type ShopInterface interface {
	CommonInterface[*modelpg.Shop]
}

type shopImpl struct{}

func Shop() ShopInterface {
	return shopImpl{}
}
