package util

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"ghttp/conf"
	"ghttp/definestruct"
	"io"
	"net/http"
	"strings"
	"time"
)

type UserStandardParam struct {
	Method  string                 `json:"method"`
	Url     string                 `json:"url"`
	Params  map[string]string      `json:"params, omitempty"`
	Body    map[string]interface{} `json:"body, omitempty"`
	Headers map[string]string      `json:"headers, omitempty"`
	Proto   string                 `json:"proto"`
	TimeOut time.Duration          `json:"time_out"`
	Cookies map[string]string      `json:"cookies, omitempty"`
}

// TODO Set proxy
// TODO Form request
// TODO Upload file
func (userParam *UserStandardParam) APIGet() *definestruct.OutStandardResult {

	userParam.CommonHandleUrl()

	uri := userParam.GetUri()

	result := userParam.baseAPI(conf.HTTP_GET, uri, nil)

	return result
}

func (userParam *UserStandardParam) APIPost() *definestruct.OutStandardResult {

	userParam.CommonHandleUrl()

	body := JsonDumps(userParam.Body)

	result := userParam.baseAPI(conf.HTTP_POST, userParam.Url, strings.NewReader(body))

	return result
}


// PUT 方法目前只支持请求体参数
func (userParam *UserStandardParam) APIPut() *definestruct.OutStandardResult {
	userParam.CommonHandleUrl()

	body := JsonDumps(userParam.Body)

	result := userParam.baseAPI(conf.HTTP_PUT, userParam.Url, strings.NewReader(body))

	return result
}

// PATCH 方法目前只支持请求体参数
func (userParam *UserStandardParam) APIPatch() *definestruct.OutStandardResult {
	userParam.CommonHandleUrl()

	body := JsonDumps(userParam.Body)

	result := userParam.baseAPI(conf.HTTP_PATCH, userParam.Url, strings.NewReader(body))

	return result
}

// DELETE 方法目前只支持请求体参数
func (userParam *UserStandardParam) APIDelete() *definestruct.OutStandardResult {
	userParam.CommonHandleUrl()

	body := JsonDumps(userParam.Body)

	result := userParam.baseAPI(conf.HTTP_DELETE, userParam.Url, strings.NewReader(body))

	return result
}

func (userParam *UserStandardParam) baseAPI(method, url string, body io.Reader) *definestruct.OutStandardResult {

	request, err := http.NewRequest(method, url, body)

	HandleErr(err, fmt.Sprintf("generate request %s failed", url))

	userParam.SetHeaders(request)

	userParam.SetCookies(request)

	newClient := NewClient(userParam.Proto, userParam.TimeOut)

	response, err := newClient.Do(request)

	HandleErr(err, fmt.Sprintf("request %s failed", url))

	defer response.Body.Close()

	result := ToStandardOut(response)

	return result
}

func (userParam *UserStandardParam) GetUri() string {

	list := make([]string, 0)
	for key, value := range userParam.Params {
		list = append(list, (key + "=" + value))
	}

	data := strings.Join(list, conf.ANOTHER_QUERY_STR_SEPARATOR)

	if strings.Contains(userParam.Url, conf.DEFAULT_QUERY_STR_SEPARATOR) {
		data = conf.ANOTHER_QUERY_STR_SEPARATOR + data
	} else {
		data = conf.DEFAULT_QUERY_STR_SEPARATOR + data
	}

	uri := userParam.Url + data

	return uri
}

func (userParam *UserStandardParam) CommonHandleUrl() {
	url := userParam.Url

	if EqualString(userParam.Proto, conf.HTTP) {
		url = conf.HTTP_SCHEMA + url
	} else {
		url = conf.HTTPS_SCHEMA + url
	}

	userParam.Url = url
}

func (userParam *UserStandardParam) SetHeaders(request *http.Request) {
	headerMap := userParam.Headers

	if _, exist := headerMap[conf.CONTENT_TYPE]; !exist {
		headerMap[conf.CONTENT_TYPE] = conf.DEFAULT_CONTENT_TYPE_VALUE
	}

	if _, exist := headerMap[conf.ACCEPT_LANGUAGE]; !exist {
		headerMap[conf.ACCEPT_LANGUAGE] = conf.DEFAULT_ACCEPT_LANGUAGE_VALUE
	}

	if _, exist := headerMap[conf.USER_AGENT]; !exist {
		headerMap[conf.USER_AGENT] = fmt.Sprintf(conf.DEFAULT_USER_AGENT_VALUE, conf.VERSION)
	}

	for key, value := range headerMap {
		request.Header.Add(key, value)
	}
}

func (userParam *UserStandardParam) SetCookies(r *http.Request) {
	cookiesMap := userParam.Cookies

	if IsEmptyStringMap(cookiesMap) {
		return
	}

	for name, value := range cookiesMap {
		cookie := &http.Cookie{
			Name:  name,
			Value: value,
		}
		r.AddCookie(cookie)
	}
}

func (userParam *UserStandardParam) ToPrettyJsonString() string {
	if userParam == nil {
		return ""
	}

	result, err := json.MarshalIndent(*userParam, "", "    ")

	if err != nil {
		return ""
	}

	return string(result)
}

func NewClient(proto string, timeout time.Duration) *http.Client {
	var client *http.Client

	if EqualString(proto, conf.HTTPS) {
		tr := &http.Transport{
			//InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true,
			//则不会校验证书以及证书中的主机名和服务器主机名是否一致。
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client = &http.Client{
			Transport: tr,
			Timeout:   timeout,
		}

	} else {
		client = &http.Client{
			Timeout: timeout,
		}
	}

	return client
}
