package definestruct

type OutStandardResult struct {
	// 状态码
	Status int `json:"status"`
	// 协议
	Proto string `json:"proto"`
	// 响应
	Response string `json:"response"`
	// 响应header
	ResponseHeaders map[string]string `json:"response_headers, omitempty"`
	// cookie
	Cookie string `json:"cookie, omitempty"`
	// 响应长度
	ContentLength int64 `json:"content_length"`
	// content-type
	ContentType string `json:"content_type"`

	Headers map[string][]string `json:"headers, omitempty"`
}
