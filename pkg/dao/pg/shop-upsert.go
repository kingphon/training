package pgdao

import (
	"context"
	"fmt"
	dbpg "go-sql/internal/pg/db"
	modelpg "go-sql/internal/pg/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (shopImpl) InsertOne(ctx context.Context, db boil.ContextExecutor, doc *modelpg.Shop) (err error) {
	if db == nil {
		db = dbpg.GetDB()
	}

	err = doc.Insert(ctx, db, boil.Infer())
	if err != nil {
		fmt.Println("error when InsertOne shop, err: ", err)
	}

	return
}
