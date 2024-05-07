package dump

import (
	utilpg "go-sql/internal/pg/util"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	Now = utilpg.SetTime(time.Now())
	DB  boil.ContextExecutor
)

func Init() {
	go dumpUser()
}
