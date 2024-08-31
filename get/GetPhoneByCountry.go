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

func GetCountryByPhone(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["phonecode"] == nil {
		return ErrorReturn(t, 500, "0000012", "Missing phonecode")
	}

	phonecode := p.Sanitize(fmt.Sprintf("%v", args["phonecode"]))

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

	db.Conn.Debug().Where("phonecode = ?", phonecode).First(&data)

	ans["country"] = data

	response = Interfacetoresponse(t, ans)
	return response
}
