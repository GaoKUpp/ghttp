package util

import (
	"errors"
	"ghttp/conf"
	"ghttp/definestruct"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	errParseBool  = errors.New("Parse bool failed")
	errParseInt   = errors.New("Parse int failed")
	errParseFloat = errors.New("Parse float failed")
	errParseSlice = errors.New("Parse slice failed")
)

func ToStandardOut(response *http.Response) *definestruct.OutStandardResult {

	body, contentLength := handleBody(response)

	result := definestruct.OutStandardResult{
		Status:        response.StatusCode,
		Proto:         response.Proto,
		ContentLength: int64(contentLength),
		ContentType:   response.Header.Get(conf.CONTENT_TYPE),
		Response:      body,
		Headers:       response.Header,
	}

	return &result
}

func handleBody(response *http.Response) (string, int) {

	body, err := ioutil.ReadAll(response.Body)

	HandleErr(err, "read body failed")

	result, err := JsonLoad(body)
	if err != nil {
		result = string(body)
	}

	return result, len(body)
}

func ToRawJsonTypeBool(v string) (bool, error) {
	if EqualString(strings.ToLower(v), strconv.FormatBool(true)) || EqualString(strings.ToLower(v), strconv.FormatBool(false)) {
		return strconv.ParseBool(v)
	}

	return false, errParseBool
}

func ToRawJsonTypeNumber(v string) (int64, error) {
	result, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		return 0, errParseInt
	}

	return result, nil
}

func ToRawJsonTypeFloat(v string) (float64, error) {
	result, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return 0.00, errParseFloat
	}

	return result, nil
}

func ToRawJsonTypeSlice(v string) ([]interface{}, error) {
	result := make([]interface{}, 0)

	if !strings.HasPrefix(v, "[") || !strings.HasSuffix(v, "]") {
		return nil, errParseSlice
	}

	v = strings.TrimPrefix(v, "[")
	v = strings.TrimSuffix(v, "]")

	strResultList := strings.Split(v, ",")

	// 将切片内部数据转化为json原生类型
	for _, ele := range strResultList {
		item := ToRawJsonType(ele)
		result = append(result, item)
	}

	return result, nil
}

func ToRawJsonType(v string) interface{} {

	if rawBoolTypeValue, err := ToRawJsonTypeBool(v); err == nil {
		return rawBoolTypeValue
	}

	if rawNumberTypeValue, err := ToRawJsonTypeNumber(v); err == nil {
		return rawNumberTypeValue
	}

	if rawFloatTypeValue, err := ToRawJsonTypeFloat(v); err == nil {
		return rawFloatTypeValue
	}

	if rawSliceTypeValue, err := ToRawJsonTypeSlice(v); err == nil {
		return rawSliceTypeValue
	}
	return v
}
