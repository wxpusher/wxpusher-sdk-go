package model

// Result 接口返回结果
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 是否为正确结果
func (r *Result) Success() bool {
	return r.Code == 1000
}
