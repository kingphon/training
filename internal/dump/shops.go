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

var ShopDumpData = modelpg.ShopSlice{
	&modelpg.Shop{
		ID:        utilpg.SetString(primitive.NewObjectID().Hex()),
		Name:      utilpg.SetString("Phương's Shop"),
		Status:    utilpg.SetString(constant.CommonStatus.Active),
		UserID:    UserDumpData[1].ID,
		Cash:      utilpg.SetFloat64(10000000),
		CreatedAt: Now,
		UpdatedAt: Now,
	},
}

func dumpShop() {
	if DB == nil {
		DB = dbpg.GetDB()
	}

	sDAO := pgdao.Shop()

	for _, u := range ShopDumpData {
		sDAO.InsertOne(context.Background(), nil, u)
	}

}
