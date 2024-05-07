package utilpg

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
)

func SetString(val string) null.String {
	return null.String{
		String: val,
		Valid:  true,
	}
}

func SetInt(val int) null.Int {
	return null.Int{
		Int:   val,
		Valid: true,
	}
}

func SetBool(val bool) (res null.Bool) {
	return null.Bool{
		Bool:  val,
		Valid: true,
	}
}

func SetTime(val time.Time) (res null.Time) {
	if val.IsZero() {
		return
	}
	return null.Time{
		Time:  val,
		Valid: true,
	}
}

func SetJSON(val []byte) (res null.JSON) {
	if val == nil {
		return
	}
	return null.JSON{
		JSON:  val,
		Valid: true,
	}
}

func SetFloat64(val float64) decimal.Decimal {
	return decimal.NewFromFloat(val)
}
