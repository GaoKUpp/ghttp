package util

import (
	"regexp"
)

type MatchResult struct {
	MatchData   string            `json:"match_data"`
	GroupResult map[string]string `json:"group_result"`
}

/*
提供优雅获取分组内容的方法
*/
func (m *MatchResult) Group(key string) string {
	return m.GroupResult[key]
}

/*
正则匹配
不匹配返回nil, false
匹配成功返回 MatchResult指针, true
*/
func Match(v, pattern string) (*MatchResult, bool) {

	// 返回与pattern符合的正则表达式
	compile := regexp.MustCompile(pattern)
	// 获取分组的命名
	groupNames := compile.SubexpNames()

	// match 是符合匹配的内容,
	// 其中第0个是匹配的数据, 后续的分组的内部的内容
	match := compile.FindStringSubmatch(v)
	if match == nil {
		return nil, false
	}

	matchedData := match[0]

	ret, err := ZipStringSlice(groupNames[1:], match[1:])

	HandleErr(err, "zip error")

	result := MatchResult{
		MatchData:   matchedData,
		GroupResult: ret,
	}

	return &result, true
}
