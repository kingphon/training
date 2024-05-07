package dump

import (
	"context"
	"go-sql/internal/constant"
	dbpg "go-sql/internal/pg/db"
	modelpg "go-sql/internal/pg/model"
	utilpg "go-sql/internal/pg/util"
	pgdao "go-sql/pkg/dao/pg"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var LocationDumpData = modelpg.LocationSlice{
	&modelpg.Location{
		ID:        utilpg.SetString(primitive.NewObjectID().Hex()),
		Name:      utilpg.SetString("Phương's Location"),
		Address:   utilpg.SetString("208 Núi Thành, Hoà Cường Bắc, Hải Châu, Đà Nẵng 550000, Việt Nam"),
		Status:    utilpg.SetString(constant.CommonStatus.Active),
		IsDefault: utilpg.SetBool(true),
		UserID:    UserDumpData[0].ID,
		Latitude:  utilpg.SetFloat64(16.044965180826427),
		Longitude: utilpg.SetFloat64(108.22118826783618),
		CreatedAt: Now,
		UpdatedAt: Now,
	},
}

func dumpLocation() {
	if DB == nil {
		DB = dbpg.GetDB()
	}

	sDAO := pgdao.Location()

	for _, u := range LocationDumpData {
		sDAO.InsertOne(context.Background(), nil, u)
	}

}
