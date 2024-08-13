package post

import (
	"dictionary/model"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func CreateCategory(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["category"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Category")
	}

	category := p.Sanitize(fmt.Sprintf("%v", args["category"]))

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	//Check does such category is exist
	curdata := &model.DictionaryCategories{}
	rows := db.Conn.Debug().Model(curdata).Where("category = ?", category).Find(&curdata)
	if rows.RowsAffected != 0 {
		return ErrorReturn(t, 500, "000005", "Such Category is exist")
	}

	data := &model.DictionaryCategories{}

	JsonArgs, err := json.Marshal(args)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = json.Unmarshal(JsonArgs, &data)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	dataid := ""

	if args["categoryid"] != nil {
		dataid = p.Sanitize(fmt.Sprintf("%v", args["categoryid"]))
	} else {
		lowerStr := strings.ToLower(p.Sanitize(fmt.Sprintf("%v", args["category"])))
		dataid = strings.ReplaceAll(lowerStr, " ", "_")
	}

	data.UUID = dataid
	data.IsCustom = true

	err = db.Conn.Create(&data).Error
	if err != nil {
		return ErrorReturn(t, 400, "000005", err.Error())
	}

	//TODO: Record event

	ans["uuid"] = dataid
	response = Interfacetoresponse(t, ans)
	return response

}
