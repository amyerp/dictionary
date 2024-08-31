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

func CreateValue(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["name"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Important Data")
	}

	categoryid := *t.Param
	value := p.Sanitize(fmt.Sprintf("%v", args["name"]))

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
	rows := db.Conn.Debug().Model(curdata).Where("uuid = ?", categoryid).Find(&curdata)
	if rows.RowsAffected == 0 {
		return ErrorReturn(t, 404, "000005", "Category not found")
	}

	curbval := &model.DictionaryValue{}
	rows = db.Conn.Debug().Model(curbval).Where("name = ? AND catrgoryid = ?", value, categoryid).Find(&curbval)
	if rows.RowsAffected != 0 {
		return ErrorReturn(t, 500, "000005", "Such Value is exist")
	}

	data := &model.DictionaryValue{}

	JsonArgs, err := json.Marshal(args)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = json.Unmarshal(JsonArgs, &data)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	lowerStr := strings.ToLower(p.Sanitize(fmt.Sprintf("%v", args["name"])))
	dataid := strings.ReplaceAll(lowerStr, " ", "_")
	if curdata.FilteredBy != "" && args["filter_value"] != nil {
		fv := strings.ToLower(p.Sanitize(fmt.Sprintf("%v", args["filter_value"])))
		dataid = fmt.Sprintf("%s_%s", dataid, fv)
	}
	data.UUID = dataid

	err = db.Conn.Create(&data).Error
	if err != nil {
		return ErrorReturn(t, 400, "000005", err.Error())
	}

	//TODO: Record event

	ans["uuid"] = dataid
	response = Interfacetoresponse(t, ans)
	return response

}
