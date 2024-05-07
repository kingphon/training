package dump

import (
	"context"
	"go-sql/internal/constant"
	dbpg "go-sql/internal/pg/db"
	modelpg "go-sql/internal/pg/model"
	utilpg "go-sql/internal/pg/util"
	pgdao "go-sql/pkg/dao/pg"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserDumpData = modelpg.UserSlice{
	&modelpg.User{
		ID:        utilpg.SetString(primitive.NewObjectID().Hex()),
		Name:      utilpg.SetString("Phương"),
		Status:    utilpg.SetString(constant.CommonStatus.Active),
		CreatedAt: Now,
		UpdatedAt: Now,
	},
	&modelpg.User{
		ID:        utilpg.SetString(primitive.NewObjectID().Hex()),
		Name:      utilpg.SetString("Phương shop"),
		Status:    utilpg.SetString(constant.CommonStatus.Active),
		CreatedAt: Now,
		UpdatedAt: Now,
	},
	&modelpg.User{
		ID:        utilpg.SetString(primitive.NewObjectID().Hex()),
		Name:      utilpg.SetString("Phương warehouse"),
		Status:    utilpg.SetString(constant.CommonStatus.Active),
		CreatedAt: Now,
		UpdatedAt: Now,
	},
}

func dumpUser() {
	if DB == nil {
		DB = dbpg.GetDB()
	}

	uDAO := pgdao.User()

	users, err := modelpg.Users().All(context.Background(), DB)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the result set is empty
	if len(users) == 0 {
		log.Println("No users found in the database")
		for _, u := range UserDumpData {
			uDAO.InsertOne(context.Background(), nil, u)
		}
	}

	dumpShop()
	dumpWarehouse()
	dumpLocation()

}
