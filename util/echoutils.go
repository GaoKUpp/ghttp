package util

import (
	"fmt"
	"ghttp/conf"
	"ghttp/definestruct"
	"strconv"
	"strings"
)

func Echo(v *definestruct.OutStandardResult) {
	if v == nil {
		return
	}

	headers := handleHeader(v.Headers)

	fmt.Println(
		Colorful(conf.BACK_BLACK, conf.FONT_GREEN, v.Proto+" " + strconv.Itoa(v.Status)),
		Colorful(conf.BACK_BLACK, conf.FONT_GREEN, "Content-Length: " + strconv.FormatInt(v.ContentLength, 10)),
		Colorful(conf.BACK_BLACK, conf.FONT_GREEN, headers),
		Colorful(conf.BACK_BLACK, conf.FONT_GREEN, "Date: " + GetUTCTimeStr()),
		"\n\n",
		Colorful(conf.BACK_BLACK, conf.FONT_PRPLE, v.Response),
	)
}

func Colorful(backEndColor conf.Color, fontColor conf.Color, content interface{}) string {
	return fmt.Sprintf("\n%c[1;%d;%dm%v%c[0m",
		0x1B,
		backEndColor,
		fontColor,
		content,
		0x1B)
}

func handleHeader(header map[string][]string) string {
	if IsEmptyMap(header) {
		return ""
	}

	_headers := make([]string, 0)

	for key, header := range header {
		data := fmt.Sprintf("%s: %s", key, header)
		_headers = append(_headers, data)
	}

	headers := strings.Join(_headers, "\n")

	return headers

}
