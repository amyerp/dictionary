package get

import (
	"dictionary/model"
	"fmt"
	"strconv"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func GetCountries(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)

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
	data := []model.Countries{}

	var count int64
	db.Conn.Debug().Model(data).Count(&count)
	db.Conn.Debug().Limit(limit).Offset(offset).Find(&data)

	ans["countries"] = data
	ans["countriescount"] = count

	response = Interfacetoresponse(t, ans)
	return response
}
