package wxpusher

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/wxpusher/wxpusher-sdk-go/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const UrlBase = "http://wxpusher.zjiecode.com"
const UrlSendMessage = UrlBase + "/api/send/message"
const UrlMessageStatus = UrlBase + "/api/send/query/${messageId}"
const UrlCreateQrcode = UrlBase + "/api/fun/create/qrcode"
const UrlQueryWxUser = UrlBase + "/api/fun/wxuser"

// SendMessage 发送消息
func SendMessage(message *model.Message) ([]model.SendMsgResult, error) {
	msgResults := make([]model.SendMsgResult, 0)
	// 校验消息内容
	err := message.Check()
	if err != nil {
		return msgResults, err
	}
	// 参数转json
	requestBody, _ := json.Marshal(message)
	resp, err := http.Post(UrlSendMessage, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	result := model.Result{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	if !result.Success() {
		return msgResults, model.NewError(result.Code, errors.New(result.Msg))
	}
	// result.Data 转为[]model.SendMsgResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &msgResults)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	return msgResults, nil
}

// QueryMessageStatus 查询消息发送状态
func QueryMessageStatus(messageId int) (model.Result, error) {
	var result model.Result
	url := strings.ReplaceAll(UrlMessageStatus, "${messageId}", strconv.Itoa(messageId))
	resp, err := http.Get(url)
	if err != nil {
		return result, model.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, model.NewSDKError(err)
	}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return result, model.NewSDKError(err)
	}
	if !result.Success() {
		return result, model.NewError(result.Code, errors.New(result.Msg))
	}
	return result, nil
}

// CreateQrcode 创建参数二维码
func CreateQrcode(qrcode *model.Qrcode) (model.CreateQrcodeResult, error) {
	var qrResult model.CreateQrcodeResult
	err := qrcode.Check()
	if err != nil {
		return qrResult, err
	}
	requestBody, _ := json.Marshal(qrcode)
	resp, err := http.Post(UrlCreateQrcode, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return qrResult, model.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return qrResult, model.NewSDKError(err)
	}
	result := model.Result{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return qrResult, model.NewSDKError(err)
	}
	if !result.Success() {
		return qrResult, model.NewError(result.Code, errors.New(result.Msg))
	}
	// result.Data 转为model.CreateQrcodeResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return qrResult, model.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &qrResult)
	if err != nil {
		return qrResult, model.NewSDKError(err)
	}
	return qrResult, nil
}

// QueryWxUser 查询App的关注用户
func QueryWxUser(appToken string, page, pageSize int) (model.QueryWxUserResult, error) {
	var queryResult model.QueryWxUserResult
	req, _ := http.NewRequest("GET", UrlQueryWxUser, nil)
	q := req.URL.Query()
	q.Add("appToken", appToken)
	q.Add("page", strconv.Itoa(page))
	q.Add("pageSize", strconv.Itoa(pageSize))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return queryResult, model.NewSDKError(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return queryResult, model.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	result := model.Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return queryResult, model.NewSDKError(err)
	}
	// result.Data 转为model.QueryWxUserResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return queryResult, model.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &queryResult)
	if err != nil {
		return queryResult, model.NewSDKError(err)
	}
	return queryResult, nil
}
