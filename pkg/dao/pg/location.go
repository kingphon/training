package pgdao

import (
	modelpg "go-sql/internal/pg/model"
)

type LocationInterface interface {
	CommonInterface[*modelpg.Location]
}

type locationImpl struct{}

func Location() LocationInterface {
	return locationImpl{}
}
