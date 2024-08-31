package get

import (
	"dictionary/model"
	"fmt"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func GetCountryByID(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["iso2"] == nil {
		return ErrorReturn(t, 500, "0000012", "Missing iso2")
	}

	iso2 := p.Sanitize(fmt.Sprintf("%v", args["iso2"]))

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

	//Show addresses
	data := model.Countries{}

	db.Conn.Debug().Where("iso2 = ?", iso2).First(&data)

	ans["country"] = data

	response = Interfacetoresponse(t, ans)
	return response
}
