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

var WarehouseDumpData = modelpg.WarehouseSlice{
	&modelpg.Warehouse{
		ID:        utilpg.SetString(primitive.NewObjectID().Hex()),
		Name:      utilpg.SetString("Phương's warehouse"),
		Address:   utilpg.SetString("WGCX+7M, Tân Phú Trung, Củ Chi, Thành phố Hồ Chí Minh, Việt Nam"),
		Status:    utilpg.SetString(constant.CommonStatus.Active),
		UserID:    UserDumpData[2].ID,
		Cash:      utilpg.SetFloat64(10000000),
		Latitude:  utilpg.SetFloat64(10.920769770987697),
		Longitude: utilpg.SetFloat64(106.5497979898543),
		CreatedAt: Now,
		UpdatedAt: Now,
	},
}

func dumpWarehouse() {
	if DB == nil {
		DB = dbpg.GetDB()
	}

	sDAO := pgdao.Warehouse()

	for _, u := range WarehouseDumpData {
		sDAO.InsertOne(context.Background(), nil, u)
	}

}
