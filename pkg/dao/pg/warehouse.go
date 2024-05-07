package pgdao

import (
	modelpg "go-sql/internal/pg/model"
)

type WarehouseInterface interface {
	CommonInterface[*modelpg.Warehouse]
}

type warehouseImpl struct{}

func Warehouse() WarehouseInterface {
	return warehouseImpl{}
}
