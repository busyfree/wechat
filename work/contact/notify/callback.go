package notify

import (
	"encoding/xml"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/contact"
)

// SignatureOptions 微信服务器验证参数
// 请求地址：回调地址/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&timestamp=13500001234&nonce=123412323&echostr=ENCRYPT_STR
type SignatureOptions struct {
	Signature string `form:"msg_signature" json:"msg_signature"`
	TimeStamp string `form:"timestamp" json:"timestamp"`
	Nonce     string `form:"nonce" json:"nonce"`
	EchoStr   string `form:"echostr" json:"echostr"`
}

// VerifyURL 验证请求参数是否合法并返回解密后的消息内容
//  //Gin框架的使用示例
//	r.GET("/v1/event/callback", func(c *gin.Context) {
//		options := kf.SignatureOptions{}
//		//获取回调的的校验参数
//		if = c.ShouldBindQuery(&options); err != nil {
//			c.String(http.StatusUnauthorized, "参数解析失败")
//		}
//		// 调用VerifyURL方法校验当前请求，如果合法则把解密后的内容作为响应返回给微信服务器
//		echo, err := kfClient.VerifyURL(options)
//		if err == nil {
//			c.String(http.StatusOK, echo)
//		} else {
//			c.String(http.StatusUnauthorized, "非法请求来源")
//		}
//	})
func (r *Notify) VerifyURL(options SignatureOptions) (string, error) {
	if options.Signature != util.Signature(r.ctx.Token, options.TimeStamp, options.Nonce, options.EchoStr) {
		return "", contact.NewSDKErr(40015)
	}
	_, bData, err := util.DecryptMsg(r.corpID, options.EchoStr, r.encodingAESKey)
	if err != nil {
		return "", contact.NewSDKErr(40016)
	}

	return string(bData), nil
}

// callbackOriginMessage 原始回调消息内容
// <xml>
// <ToUserName><![CDATA[toUser]]></ToUserName>
// <AgentID><![CDATA[toAgentID]]></AgentID>
// <Encrypt><![CDATA[msg_encrypt]]></Encrypt>
// </xml>
type callbackOriginMessage struct {
	ToUserName string `xml:"ToUserName"` // 企业微信的CorpID，当为第三方套件回调事件时，CorpID的内容为suiteid
	AgentID    string `xml:"AgentID"`    // 接收的应用id，可在应用的设置页面获取
	Encrypt    string `xml:"Encrypt"`    // 消息结构体加密后的字符串
}

// CallbackMessageResp 通讯录回调响应消息
// <xml>
// <Encrypt><![CDATA[msg_encrypt]]></Encrypt>
// <MsgSignature><![CDATA[msg_signature]]></MsgSignature>
// <TimeStamp>timestamp</TimeStamp>
// <Nonce><![CDATA[nonce]]></Nonce>
// </xml>
type CallbackMessageResp struct {
	Encrypt      string `json:"Encrypt" xml:"Encrypt"`           // 微信客服组件ID
	MsgSignature int    `json:"MsgSignature" xml:"MsgSignature"` // 消息创建时间，unix时间戳
	TimeStamp    string `json:"TimeStamp" xml:"TimeStamp"`       // 消息的类型，此时固定为 event
	Nonce        string `json:"Nonce" xml:"Nonce"`               // 事件的类型，此时固定为 kf_msg_or_event
}

// GetCallbackRawMsg 获取回调事件中的消息内容
//  //Gin框架的使用示例
//	r.POST("/v1/event/callback", func(c *gin.Context) {
//		var (
//			message kf.CallbackMessage
//			body []byte
//		)
//		// 读取原始消息内容
//		body, err = c.GetRawData()
//		if err != nil {
//			c.String(http.StatusInternalServerError, err.Error())
//			return
//		}
//		// 解析原始数据
//		message, err = kfClient.GetCallbackMessage(body)
//		if err != nil {
//			c.String(http.StatusInternalServerError, "消息获取失败")
//			return
//		}
//		fmt.Println(message)
//		c.String(200, "ok")
//	})
func (r *Notify) GetCallbackRawMsg(encryptedMsg []byte) (msg []byte, err error) {
	var origin callbackOriginMessage
	if err = xml.Unmarshal(encryptedMsg, &origin); err != nil {
		return msg, err
	}
	_, msg, err = util.DecryptMsg(r.corpID, origin.Encrypt, r.encodingAESKey)
	if err != nil {
		err = contact.NewSDKErr(40016)
		return
	}
	return
	// if err = xml.Unmarshal(bData, &msg); err != nil {
	// 	return msg, err
	// }
	// return msg, err
}
