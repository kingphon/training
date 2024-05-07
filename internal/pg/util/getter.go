package utilpg

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
)

func GetString(val null.String) string {
	if !val.Valid {
		return ""
	}
	return val.String
}

func GetInt(val null.Int) int {
	if !val.Valid {
		return 0
	}
	return val.Int
}

func GetBool(val null.Bool) bool {
	if !val.Valid {
		return false
	}
	return val.Bool
}

func GetTime(val null.Time) (res time.Time) {
	if !val.Valid {
		return
	}
	return val.Time
}

func GetJSON(val null.JSON) (res []byte) {
	if !val.Valid {
		return
	}
	return val.JSON
}

// GetFloat64 ...
func GetFloat64(val decimal.Decimal) float64 {
	return val.InexactFloat64()
}
