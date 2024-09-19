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

func GetValueByID(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()
	categoryid := ""

	if args["uuid"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  UUID")
	}

	uuid := p.Sanitize(fmt.Sprintf("%v", args["uuid"]))

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

	if args["categoryid"] != nil {
		categoryid = p.Sanitize(fmt.Sprintf("%v", args["categoryid"]))
	}

	//Show addresses
	data := &model.DictionaryValue{}

	if categoryid != "" {

		curdata := &model.DictionaryCategories{}
		rows := db.Conn.Debug().Model(curdata).Where("category = ? OR uuid = ?", categoryid, categoryid).Find(&curdata)
		if rows.RowsAffected == 0 {
			return ErrorReturn(t, 404, "000005", "Such Category is not exist")
		}
		catid := curdata.UUID

		db.Conn.Debug().Where("categoryid = ? AND uuid = ?", catid, uuid).First(&data)
	} else {
		db.Conn.Debug().Where("uuid = ?", uuid).First(&data)
	}

	ans["dictionary"] = data

	response = Interfacetoresponse(t, ans)
	return response
}
