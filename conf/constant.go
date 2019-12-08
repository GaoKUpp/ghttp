package conf

const VERSION = "1.0"

// SUPPORTED HTTP METHOD
const (
	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_PATCH  = "PATCH"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"
)

// HTTP COMMON CONSTANT
const (
	HTTP  = "http"
	HTTPS = "https"

	HTTP_SCHEMA                 = "http://"
	HTTPS_SCHEMA                = "https://"
	FINAL_SIGN                  = "/"
	DEFAULT_QUERY_STR_SEPARATOR = "?"
	ANOTHER_QUERY_STR_SEPARATOR = "&"

	CONTENT_TYPE               = "Content-Type"
	DEFAULT_CONTENT_TYPE_VALUE = "application/json;charset=UTF-8"

	ACCEPT_LANGUAGE               = "Accept-Language"
	DEFAULT_ACCEPT_LANGUAGE_VALUE = "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"

	COOKIE_SIGN      = "Cookie"
	COOKIE_SEPARATOR = ";"

	USER_AGENT               = "User-Agent"
	DEFAULT_USER_AGENT_VALUE = "G_HTTPie/%s"

	DEFAULT_TIME_OUT = 0
)

// parse sign
const (
	PARAM_LEAST_NUMBER = 2

	QUERY_STR_SIGN    = "=="
	RE_REQUEST_BODY   = "(?P<key>\\w+?)(?P<sign>:?=)(?P<value>[^=].*)"
	RE_REQUEST_HEADER = "(?P<key>\\w+?):(?P<value>[^=].*)"
	RAW_JSON_SIGN     = ":="
	REQUEST_BODY_SIGN = "="
)

type Color int

// 背景色
const (
	BACK_BLACK Color = iota + 39
	BACK_WHITE
	BACK_RED
	BACK_GREEN
	BACK_BROWN
	BACK_PRPLE
	BACK_BLUE
	BACK_GRAY
)

// 字体颜色
const (
	FONT_WHITE Color = iota + 30
	FONT_RED
	FONT_GREEN
	FONT_BROWN
	FONT_PRPLE
	FONT_PRPLE_AND_RED
	FONT_BLUE
	FONT_GRAY
)
