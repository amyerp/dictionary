package get

import (
	"dictionary/model"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

// api/v3/dictionary/{category}/
func Init(t *pb.Request) (response *pb.Response) {
	switch *t.Param {
	case "country":
		response = GetCountries(t)
	case "phone_codes":
		response = GetPhoneCodes(t)
	case "currency":
		response = GetCurrencies(t)
	case "cities":
		response = GetCities(t)
	case "states":
		response = GetStates(t)
	case "categories":
		response = GetCategories(t)
	default:
		response = CheckCategoryID(t)
	}

	return response
}

func CheckCategoryID(t *pb.Request) (response *pb.Response) {
	category := *t.Param

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	curdata := &model.DictionaryCategories{}
	rows := db.Conn.Debug().Model(curdata).Where("category = ? OR uuid = ?", category, category).Find(&curdata)
	if rows.RowsAffected == 0 {
		return ErrorReturn(t, 404, "000005", "Such Category is not exist")
	}
	return GetValuesByCategory(t)
}
