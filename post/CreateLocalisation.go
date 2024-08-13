package post

import (
	"dictionary/model"
	"encoding/json"
	"fmt"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func CreateLocalisation(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["valueid"] == nil || args["value"] == nil || args["language"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Important Data")
	}

	valueid := p.Sanitize(fmt.Sprintf("%v", args["valueid"]))
	value := p.Sanitize(fmt.Sprintf("%v", args["value"]))
	language := p.Sanitize(fmt.Sprintf("%v", args["language"]))

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	curbval := &model.DictionaryValue{}
	rows := db.Conn.Debug().Model(curbval).Where("uuid = ? ", valueid).Find(&curbval)
	if rows.RowsAffected == 0 {
		return ErrorReturn(t, 404, "000005", "Value not found")
	}

	curlang := &model.DictionaryValueLoc{}
	rows = db.Conn.Debug().Model(curlang).Where("value = ? AND  language=?", value, language).Find(&curlang)
	if rows.RowsAffected != 0 {
		return ErrorReturn(t, 500, "000005", "Such localisation is exist")
	}

	data := &model.DictionaryValueLoc{}

	JsonArgs, err := json.Marshal(args)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = json.Unmarshal(JsonArgs, &data)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = db.Conn.Create(&data).Error
	if err != nil {
		return ErrorReturn(t, 400, "000005", err.Error())
	}

	//TODO: Record event

	ans["valueid"] = valueid
	response = Interfacetoresponse(t, ans)
	return response

}
