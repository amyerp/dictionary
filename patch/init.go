package patch

import (
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {
	switch *t.Param {
	case "category":
		response = PatchCategory(t)
	default:
		response = CheckCategoryID(t)
	}

	return response
}

func CheckCategoryID(t *pb.Request) (response *pb.Response) {
	switch *t.ParamID {
	case "value":
		response = PatchValue(t)
	default:
		response = CheckLocalisation(t)
	}

	return response
}

func CheckLocalisation(t *pb.Request) (response *pb.Response) {
	switch *t.ParamIDD {
	case "loc":
		response = PatchLocalisation(t)
	default:
		response = ErrorReturn(t, 404, "00004", "Wrong request")
	}

	return response
}
