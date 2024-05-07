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

func (userImpl) FindRawByID(ctx context.Context, db boil.ContextExecutor, id string) (result *modelpg.User, err error) {
	if db == nil {
		db = dbpg.GetDB()
	}

	res := builderpg.User().FindRawByID(ctx, id)
	if res.Err != nil {
		fmt.Println("error when build user FindRawByID SQL, err: ", err)
		return &modelpg.User{}, err
	}

	result, err = modelpg.Users(qm.SQL(res.SQL, res.Args...)).One(ctx, db)

	if err != nil {
		fmt.Println("error when FindRawByID, err: ", err)
	}

	return
}

// func (userImpl) FindRawsByCondition(ctx context.Context, db boil.ContextExecutor, options optionmodel.Options) (result modelpg.UserSlice, err error) {
// 	if db == nil {
// 		db = dbpg.GetDB()
// 	}
//
// 	res := builderpg.User().FindRawsByCondition(ctx, options)
// 	if res.Err != nil {
// 		fmt.Println("error when build user FindRawsByCondition SQL, err: ", err)
// 		return nil, err
// 	}
//
// 	err = modelpg.Users(qm.SQL(res.SQL, res.Args...)).Bind(ctx, db, result)
//
// 	if err != nil {
// 		fmt.Println("error when FindRawsByCondition, err: ", err)
// 	}
//
// 	return
// }
