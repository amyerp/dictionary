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

func GetValuesByCategory(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()
	categoryid := ""

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
	} else {
		//	return ErrorReturn(t, 406, "000012", "Missing  Category ID")
		//We should have GetValuesByCategory
		categoryid = *t.Param

	}

	curdata := &model.DictionaryCategories{}
	rows := db.Conn.Debug().Model(curdata).Where("category = ? OR uuid = ?", categoryid, categoryid).Find(&curdata)
	if rows.RowsAffected == 0 {
		return ErrorReturn(t, 404, "000005", "Such Category is not exist")
	}
	catid := curdata.UUID

	//Sorting
	offset := 0
	limit := 25
	filter := ""

	if args["offset"] != nil {
		offset, _ = strconv.Atoi(fmt.Sprintf("%v", args["offset"]))
	}

	if args["limit"] != nil {
		limit, _ = strconv.Atoi(fmt.Sprintf("%v", args["limit"]))
	}
	/*
		if curdata.FilteredBy != "" && args["filter"] == nil {
			return ErrorReturn(t, 404, "000005", "Missing Filter")
		}
	*/
	if args["filter"] != nil {
		filter = p.Sanitize(fmt.Sprintf("%v", args["filter"]))
	}

	//Show addresses
	data := []model.DictionaryValue{}

	var count int64
	if filter != "" {
		db.Conn.Debug().Model(data).Where("catrgoryid = ? AND (filter_value = ? OR filter_value = ?)", catid, filter, "all").Count(&count)
		db.Conn.Debug().Where("catrgoryid = ?  AND (filter_value = ? OR filter_value = ?)", catid, filter, "all").Limit(limit).Offset(offset).Find(&data)
	} else {
		db.Conn.Debug().Model(data).Where("catrgoryid = ?", catid).Count(&count)
		db.Conn.Debug().Where("catrgoryid = ?", catid).Limit(limit).Offset(offset).Find(&data)
	}

	ans["dictionary"] = data
	ans["count"] = count

	response = Interfacetoresponse(t, ans)
	return response
}
