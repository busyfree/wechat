package notify

import (
	"encoding/xml"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/xerror"
)

// SignatureOptions 微信服务器验证参数
//
// 请求地址：url?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&timestamp=13500001234&nonce=123412323&echostr=ENCRYPT_STR
type SignatureOptions struct {
	Signature string `form:"msg_signature" json:"msg_signature"`
	TimeStamp string `form:"timestamp" json:"timestamp"`
	Nonce     string `form:"nonce" json:"nonce"`
	EchoStr   string `form:"echostr" json:"echostr"`
}

// VerifyURL 验证请求参数是否合法并返回解密后的消息内容
func (r *Notify) VerifyURL(options SignatureOptions) (string, error) {
	if options.Signature != util.Signature(r.ctx.Token, options.TimeStamp, options.Nonce, options.EchoStr) {
		return "", xerror.NewSDKErr(40015)
	}
	_, bData, err := util.DecryptMsg(r.corpID, options.EchoStr, r.encodingAESKey)
	if err != nil {
		return "", xerror.NewSDKErr(40016)
	}

	return string(bData), nil
}

// callbackOriginMessage 原始回调消息内容
//
// <xml>
//		<ToUserName><![CDATA[toUser]]></ToUserName>
//		<AgentID><![CDATA[toAgentID]]></AgentID>
//		<Encrypt><![CDATA[msg_encrypt]]></Encrypt>
// </xml>
type callbackOriginMessage struct {
	ToUserName string `xml:"ToUserName"` // 企业微信的CorpID，当为第三方套件回调事件时，CorpID的内容为suiteid
	AgentID    string `xml:"AgentID"`    // 接收的应用id，可在应用的设置页面获取
	Encrypt    string `xml:"Encrypt"`    // 消息结构体加密后的字符串
}

// GetCallbackMsg 获取回调事件中的消息内容
func (r *Notify) GetCallbackMsg(encryptedRawMsg []byte) (plainTxtByte []byte, msg MixedMsg, err error) {
	var origin callbackOriginMessage
	if err = xml.Unmarshal(encryptedRawMsg, &origin); err != nil {
		return
	}
	_, plainTxtByte, err = util.DecryptMsg(r.corpID, origin.Encrypt, r.encodingAESKey)
	if err != nil {
		err = xerror.NewSDKErr(40016)
		return
	}
	if err = xml.Unmarshal(plainTxtByte, &msg); err != nil {
		return
	}
	return
}
