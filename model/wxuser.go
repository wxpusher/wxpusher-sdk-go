package model

// WxUser 微信用户
type WxUser struct {
	UId        string `json:"uid"`
	Enable     bool   `json:"enable"`
	CreateTime int64  `json:"createTime"`
	NickName   string `json:"nickName"`
	HeadImg    string `json:"headImg"`
}

// QueryWxUserResult 分页查询微信用户返回结果
type QueryWxUserResult struct {
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
	Records  []WxUser `json:"records"`
}
