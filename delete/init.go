package delete

import (
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {
	switch *t.Param {
	case "category":
		response = DelCategory(t)
	default:
		response = CheckCategoryID(t)
	}

	return response
}

func CheckCategoryID(t *pb.Request) (response *pb.Response) {
	switch *t.ParamID {
	case "value":
		response = DelCategory(t)
	default:
		response = CheckLocalisation(t)
	}

	return response
}

func CheckLocalisation(t *pb.Request) (response *pb.Response) {
	switch *t.ParamIDD {
	case "loc":
		response = DelCategory(t)
	default:
		response = ErrorReturn(t, 404, "00004", "Wrong request")
	}

	return response
}
