package command

import (
	"fmt"
	"ghttp/conf"
	"ghttp/definestruct"
	"ghttp/util"
	"log"
)

func Request(userStandParam *util.UserStandardParam) *definestruct.OutStandardResult {

	switch userStandParam.Method {

	case conf.HTTP_GET:
		return userStandParam.APIGet()

	case conf.HTTP_POST:
		return userStandParam.APIPost()

	case conf.HTTP_PUT:
		return userStandParam.APIPut()

	case conf.HTTP_PATCH:
		return userStandParam.APIPatch()

	case conf.HTTP_DELETE:
		return userStandParam.APIDelete()

	default:
		log.Fatal(fmt.Sprintf("Unsupported method: %s", userStandParam.Method))
	}

	return nil
}
