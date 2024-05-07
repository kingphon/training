package pgdao

import (
	"context"
	"fmt"
	dbpg "go-sql/internal/pg/db"
	modelpg "go-sql/internal/pg/model"
	builderpg "go-sql/pkg/dao/pg/builder"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (warehouseImpl) FindRawByID(ctx context.Context, db boil.ContextExecutor, id string) (result *modelpg.Warehouse, err error) {
	if db == nil {
		db = dbpg.GetDB()
	}

	res := builderpg.Warehouse().FindRawByID(ctx, id)
	if res.Err != nil {
		fmt.Println("error when build warehouse FindRawByID SQL, err: ", err)
		return &modelpg.Warehouse{}, err
	}

	result, err = modelpg.Warehouses(qm.SQL(res.SQL, res.Args...)).One(ctx, db)

	if err != nil {
		fmt.Println("error when FindRawByID, err: ", err)
	}

	return
}
