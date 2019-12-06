package command

import (
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
	default:
		log.Fatal("error")
	}

	return nil
}
