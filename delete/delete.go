package delete

import (
	"dictionary/model"
	"fmt"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func DelCategory(t *pb.Request) (response *pb.Response) {
	p := bluemonday.UGCPolicy()
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}
		return ErrorReturn(t, 500, "000027", err.Error())

	}

	if p.Sanitize(fmt.Sprintf("%v", args["categoryid"])) != "" {
		id := p.Sanitize(fmt.Sprintf("%v", args["categoryid"]))
		db.Conn.Delete(model.DictionaryCategories{}, "uuid = ?", id)
		ans["response"] = "200501" // Business deleted
		response = Interfacetoresponse(t, ans)
		return response
	}

	if p.Sanitize(fmt.Sprintf("%v", args["valueid"])) != "" {
		id := p.Sanitize(fmt.Sprintf("%v", args["valueid"]))
		if p.Sanitize(fmt.Sprintf("%v", args["language"])) != "" {
			language := p.Sanitize(fmt.Sprintf("%v", args["language"]))
			db.Conn.Delete(model.DictionaryValueLoc{}, "valueid = ? AND language = ?", id, language)
		} else {
			db.Conn.Delete(model.DictionaryValue{}, "uuid = ?", id)
		}

		ans["response"] = "200501" // Business deleted
		response = Interfacetoresponse(t, ans)
		return response
	}

	return response

}
