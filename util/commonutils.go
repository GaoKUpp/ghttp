package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"ghttp/conf"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	lengthErr = errors.New("invalid arguments: two slice length not same")
)

/*
仿照python 的zip函数
将两个切片合并程一个map
目前只支持字符串切片
*/
func ZipStringSlice(keySlice, valueSlice []string) (map[string]string, error) {
	if len(keySlice) != len(valueSlice) {
		return nil, lengthErr
	}

	result := make(map[string]string)

	if keySlice == nil || valueSlice == nil {
		return result, nil
	}

	length := len(keySlice)

	for i := 0; i < length; i++ {
		result[keySlice[i]] = valueSlice[i]
	}

	return result, nil
}

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("err: %v, msg: %s", err, msg))
	}
}

func JsonDumps(v interface{}) string {
	data, err := json.Marshal(v)

	if err != nil {
		log.Fatal(fmt.Sprintf("parse json failed: %v\n", v))
	}

	return string(data)
}

func JsonLoad(v []byte) (string, error) {

	r := make(map[string]interface{})

	err := json.Unmarshal(v, &r)

	if err != nil {
		return "", err
	}

	result, _ := json.MarshalIndent(r, "", "    ")

	return string(result), nil
}

func IsQueryStr(v string) bool {

	result := false
	if strings.Contains(v, conf.QUERY_STR_SIGN) {
		result = true
	}

	return result
}

func GetUTCTimeStr() string {
	now := time.Now().UTC()
	// 显示时间格式： UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	return fmt.Sprintf("%s", now.Format(time.UnixDate))
}

func ReadStdin() map[string]interface{} {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]interface{})

	if info.Mode() & os.ModeCharDevice != 0 || info.Size() <= 0 {
		return result
	}

	reader := bufio.NewReader(os.Stdin)
	output := make([]byte, 0)

	for {
		input, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	_ = json.Unmarshal(output, &result)

	return result
}
