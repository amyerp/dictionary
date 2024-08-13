package main

import (
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"

	ad "dictionary/admin"
	dl "dictionary/delete"
	gt "dictionary/get"
	patch "dictionary/patch"
	pt "dictionary/post"
	. "dictionary/version"
)

func Init(t *pb.Request) (response *pb.Response) {

	switch *t.Param {
	case "admin":
		return admincheck(t)
	}

	if *t.Method == "GET" {

		switch *t.Param {
		case "info":
			response = info(t)
		case "health":
			response = health(t)
		default:
			response = gt.Init(t)
		}
	} else {
		if t.UID == nil {
			response = ErrorReturn(t, 401, "00002", "You are not authorized")
		}

		switch *t.Method {
		case "POST":
			response = pt.Init(t)
		case "PATCH":
			response = patch.Init(t)
		case "DELETE":
			response = dl.Init(t)
		default:
			response = ErrorReturn(t, 404, "00004", "Wrong Method")

		}

	}

	return response

}

func info(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	ans["pluginname"] = "dictionary"
	ans["version"] = VERSIONPLUGIN
	ans["description"] = ""
	response = Interfacetoresponse(t, ans)
	return response
}

func health(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	ans["health"] = "OK"
	response = Interfacetoresponse(t, ans)
	return response
}

func admincheck(t *pb.Request) (response *pb.Response) {

	if *t.IsAdmin != 1 {
		response = ErrorReturn(t, 401, "000012", "You have no admin rights")
	}

	return ad.Init(t)

}
