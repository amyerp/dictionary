package get

import (
	"dictionary/model"
	"fmt"
	"strconv"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func GetStates(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["country"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Country ID")
	}

	countryid := p.Sanitize(fmt.Sprintf("%v", args["country"]))

	//Check DB and table config
	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}
		return ErrorReturn(t, 500, "000027", err.Error())
	}

	//Sorting
	offset := 0
	limit := 25

	if args["offset"] != nil {
		offset, _ = strconv.Atoi(fmt.Sprintf("%v", args["offset"]))
	}

	if args["limit"] != nil {
		limit, _ = strconv.Atoi(fmt.Sprintf("%v", args["limit"]))
	}

	//Show addresses
	data := []model.States{}

	var count int64
	db.Conn.Debug().Model(data).Where("country_code = ?", countryid).Count(&count)
	db.Conn.Debug().Where("country_code = ?", countryid).Limit(limit).Offset(offset).Find(&data)

	ans["states"] = data
	ans["statescount"] = count

	response = Interfacetoresponse(t, ans)
	return response
}
