//////////////////////////////////////////////////////////////////////////////////
// Copyright 2021-2024 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package patch

import (
	"encoding/json"
	"fmt"

	"dictionary/model"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"

	"github.com/microcosm-cc/bluemonday"
)

func PatchLocalisation(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["valueid"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Value ID")
	}

	dataid := p.Sanitize(fmt.Sprintf("%v", args["valueid"]))

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
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

	err = db.Conn.Where("valueid = ?", dataid).Updates(&data).Error
	if err != nil {
		return ErrorReturn(t, 400, "000005", err.Error())
	}

	//TODO: Record event

	ans["uuid"] = dataid
	response = Interfacetoresponse(t, ans)
	return response

}
