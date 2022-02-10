package jssdk

import (
	"fmt"

	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

// JsSDK struct
type JsSDK struct {
	*context.Context
	credential.JsTicketHandle
}

// Config 返回给用户jssdk配置信息
type Config struct {
	CorpID    string `json:"corpid"`
	AgentID   string `json:"agentid,omitempty"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

// NewJsSDK init
func NewJsSDK(context *context.Context) *JsSDK {
	js := new(JsSDK)
	js.Context = context
	jsTicketHandle := credential.NewDefaultJsTicket(context.CorpID, credential.CacheKeyWorkPrefix, context.Cache)
	js.SetJsTicketHandle(jsTicketHandle)
	return js
}

// SetJsTicketHandle 自定义js ticket取值方式
func (js *JsSDK) SetJsTicketHandle(ticketHandle credential.JsTicketHandle) {
	js.JsTicketHandle = ticketHandle
}

// GetCorpConfig 获取jssdk需要的配置参数
// uri 为当前网页地址
func (js *JsSDK) GetCorpConfig(uri string) (config *Config, err error) {
	config = new(Config)
	var accessToken string
	accessToken, err = js.GetAccessToken()
	if err != nil {
		return
	}
	var ticketStr string
	ticketStr, err = js.GetTicket(accessToken, "work")
	if err != nil {
		return
	}

	nonceStr := util.RandomStr(16)
	timestamp := util.GetCurrTS()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, uri)
	sigStr := util.Signature(str)

	config.CorpID = js.CorpID
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}

// GetCorpAgentConfig 获取jssdk需要的配置参数
// uri 为当前网页地址
func (js *JsSDK) GetCorpAgentConfig(uri string) (config *Config, err error) {
	config = new(Config)
	var accessToken string
	accessToken, err = js.GetAccessToken()
	if err != nil {
		return
	}
	var ticketStr string
	ticketStr, err = js.GetTicket(accessToken, "agent")
	if err != nil {
		return
	}

	nonceStr := util.RandomStr(16)
	timestamp := util.GetCurrTS()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, uri)
	sigStr := util.Signature(str)

	config.CorpID = js.CorpID
	config.AgentID = js.AgentID
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}
