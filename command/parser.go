package command

import (
	"flag"
	"ghttp/conf"
	"ghttp/util"
	"log"
	"strings"
	"time"
)

var (
	Debug    bool
	Timeout  int
	UseHttps bool

	userInput      []string
	UserStandParam util.UserStandardParam
)

func parseUserInputEle() {
	allElements := flag.Args()

	if len(allElements) < conf.PARAM_LEAST_NUMBER {
		log.Fatal(conf.INVALID_PARAMS)
	}

	userInput = allElements
}

func parseMethod() string {
	method := userInput[0]

	switch strings.ToUpper(method) {
	case conf.HTTP_GET:
		method = conf.HTTP_GET
	case conf.HTTP_POST:
		method = conf.HTTP_POST
	case conf.HTTP_PATCH:
		method = conf.HTTP_PATCH
	case conf.HTTP_PUT:
		method = conf.HTTP_PUT
	case conf.HTTP_DELETE:
		method = conf.HTTP_DELETE
	default:
		log.Fatal(conf.INVALID_HTTP_METHOD)
	}

	return method
}

func parseUrl() string {
	url := userInput[1]

	if strings.HasPrefix(url, conf.HTTP_SCHEMA) {
		url = strings.TrimPrefix(url, conf.HTTP_SCHEMA)
	} else if strings.HasPrefix(url, conf.HTTPS_SCHEMA) {
		url = strings.TrimPrefix(url, conf.HTTPS_SCHEMA)
	}

	return url
}

func parseQueryStr() map[string]string {
	notVerifyEle := userInput[conf.PARAM_LEAST_NUMBER:]
	result := make(map[string]string)

	for _, ele := range notVerifyEle {
		if util.IsQueryStr(ele) {
			tmpList := strings.Split(ele, conf.QUERY_STR_SIGN)
			result[tmpList[0]] = tmpList[1]
		}
	}

	return result
}

func parseBody() map[string]interface{} {
	notVerifyEle := userInput[conf.PARAM_LEAST_NUMBER:]
	// 读取管道符输出的内容
	result := util.ReadStdin()

	for _, ele := range notVerifyEle {
		// 目前正则存在bug，原本意图是匹配 :=和=这两种
		// 但是目前这种正则会将:==这种正常的数据情况丢失,
		// 像(aa)(:=)(=1122)这样的数据, 无法匹配成功
		data, ok := util.Match(ele, conf.RE_REQUEST_BODY)

		if ok {
			sign := data.Group("sign")
			// 目前这种类型转化对于发起请求不会有影响, 考虑将来是不是可将其移除
			if util.EqualString(sign, conf.REQUEST_BODY_SIGN) {
				result[data.Group("key")] = data.Group("value")
			} else {
				result[data.Group("key")] = util.ToRawJsonType(data.Group("value"))
			}

		}
	}

	return result
}

func parseHeader() map[string]string {
	notVerifyEle := userInput[conf.PARAM_LEAST_NUMBER:]

	result := make(map[string]string)

	for _, ele := range notVerifyEle {
		matchResult, ok := util.Match(ele, conf.RE_REQUEST_HEADER)
		if ok {
			result[matchResult.Group("key")] = matchResult.Group("value")
		}
	}

	return result
}

func parseCookie() map[string]string {

	result := make(map[string]string)

	cookieStr, ok := UserStandParam.Headers[conf.COOKIE_SIGN]

	if ok {
		cookies := strings.Split(cookieStr, conf.COOKIE_SEPARATOR)

		for _, item := range cookies {
			tmp := strings.Split(item, "=")
			result[tmp[0]] = tmp[1]
		}

		delete(UserStandParam.Headers, conf.COOKIE_SIGN)
	}

	return result
}

func parseProto() string {
	result := conf.HTTP
	if UseHttps {
		result = conf.HTTPS
	}
	return result
}

func parseTimeOut() time.Duration {
	return time.Duration(Timeout) * time.Second
}

func parseMustEle() {
	UserStandParam.Method = parseMethod()
	UserStandParam.Url = parseUrl()
}

func parseOptionEle() {
	if util.EqualString(UserStandParam.Method, conf.HTTP_GET) {
		UserStandParam.Params = parseQueryStr()
	} else {
		UserStandParam.Body = parseBody()
	}
	UserStandParam.Headers = parseHeader()
	UserStandParam.Cookies = parseCookie()
	UserStandParam.Proto = parseProto()
	UserStandParam.TimeOut = parseTimeOut()
}

func Parse() *util.UserStandardParam {
	flag.Parse()

	//fmt.Printf("debug: %t, Timeout: %d, UseHttps: %t\n", Debug, Timeout, UseHttps)

	parseUserInputEle()

	parseMustEle()

	parseOptionEle()

	return &UserStandParam
}

func init() {
	flag.BoolVar(&Debug, "d", false, "Debug 标志")
	flag.IntVar(&Timeout, "t", conf.DEFAULT_TIME_OUT, "超时时间")
	flag.BoolVar(&UseHttps, "s", false, "协议")
}
