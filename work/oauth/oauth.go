package oauth

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

// Oauth auth
type Oauth struct {
	*context.Context
}

var (
	// oauthTargetURL 企业微信内跳转地址
	oauthTargetURL    = "/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
	oauthTargetWebURL = "/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=%s&agentid=%d#wechat_redirect"
	// oauthUserInfoURL 获取用户信息地址
	oauthUserInfoURL          = "/cgi-bin/user/getuserinfo?access_token=%s&code=%s"
	oauthGetUserInfoURL       = "/cgi-bin/auth/getuserinfo?access_token=%s&code=%s"
	oauthGetUserDetailInfoURL = "/cgi-bin/auth/getuserdetail?access_token=%s"
	// oauthQrContentTargetURL 构造独立窗口登录二维码
	oauthQrContentTargetURL = "/wwopen/sso/qrConnect?appid=%s&agentid=%d&redirect_uri=%s&state=%s"
)

// NewOauth new init oauth
func NewOauth(ctx *context.Context) *Oauth {
	return &Oauth{
		ctx,
	}
}

// GetTargetURL 获取授权地址
func (ctr *Oauth) GetTargetURL(callbackURL string) string {
	// url encode
	return ctr.GetOpenAPIDomain() + fmt.Sprintf(
		oauthTargetURL,
		ctr.CorpID,
		url.QueryEscape(callbackURL),
	)
}

// GetWebTargetURL 获取授权地址
func (ctr *Oauth) GetWebTargetURL(callbackURL string, agentId int) (string, string) {
	// url encode
	state := util.RandomStr(16)
	return ctr.GetOpenAPIDomain() + fmt.Sprintf(
		oauthTargetWebURL,
		ctr.CorpID,
		url.QueryEscape(callbackURL),
		state,
		agentId,
	), state
}

// GetQrContentTargetURL 构造独立窗口登录二维码
func (ctr *Oauth) GetQrContentTargetURL(callbackURL string) (string, string) {
	// url encode
	state := util.RandomStr(16)
	return ctr.GetOpenQYAPIDomain() + fmt.Sprintf(
		oauthQrContentTargetURL,
		ctr.CorpID,
		ctr.AgentID,
		url.QueryEscape(callbackURL),
		state,
	), state
}

// ResUserInfo 返回得用户信息
type ResUserInfo struct {
	util.CommonError
	// 当用户为企业成员时返回
	UserID   string `json:"UserId"`
	DeviceID string `json:"DeviceId"`
	// 非企业成员授权时返回
	OpenID         string `json:"OpenId"`
	ExternalUserID string `json:"external_userid"`
}

// UserFromCode 根据code获取用户信息
func (ctr *Oauth) UserFromCode(code string) (result ResUserInfo, err error) {
	var accessToken string
	if accessToken, err = ctr.GetAccessToken(); err != nil {
		return
	}
	var response []byte
	if response, err = util.HTTPGet(ctr.GetQYAPIDomain() + fmt.Sprintf(oauthUserInfoURL, accessToken, code)); err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		err = fmt.Errorf("UserFromCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

type GetUserInfoResp struct {
	util.CommonError
	UserID         string `json:"userid,omitempty"`
	OpenID         string `json:"openid,omitempty"`
	UserTicket     string `json:"user_ticket,omitempty"`
	ExternalUserID string `json:"external_userid,omitempty"`
}

// GetUserInfo 获取访问用户身份
// 该接口用于根据code获取成员信息，适用于自建应用与代开发应用
// https://developer.work.weixin.qq.com/document/path/91023
func (ctr *Oauth) GetUserInfo(code string) (result GetUserInfoResp, err error) {
	var accessToken string
	if accessToken, err = ctr.GetAccessToken(); err != nil {
		return
	}
	var response []byte
	if response, err = util.HTTPGet(ctr.GetQYAPIDomain() + fmt.Sprintf(oauthGetUserInfoURL, accessToken, code)); err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

type GetUserInfoDetailReq struct {
	UserTicket string `json:"user_ticket"`
}
type GetUserInfoDetailResp struct {
	util.CommonError
	UserID  string `json:"userid"`
	Gender  string `json:"gender"`
	Avatar  string `json:"avatar"`
	QrCode  string `json:"qr_code"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	BizMail string `json:"biz_mail"`
	Address string `json:"address"`
}

// GetUserInfoDetail 获取访问用户敏感信息
// 自建应用与代开发应用可通过该接口获取成员授权的敏感字段
// https://developer.work.weixin.qq.com/document/path/95833
func (ctr *Oauth) GetUserInfoDetail(userTicket string) (result GetUserInfoDetailResp, err error) {
	if len(userTicket) == 0 {
		err = fmt.Errorf("missing user_ticket")
		return
	}
	options := GetUserInfoDetailReq{UserTicket: userTicket}
	var accessToken string
	if accessToken, err = ctr.GetAccessToken(); err != nil {
		return
	}
	var response []byte
	if response, err = util.PostJSON(ctr.GetQYAPIDomain()+fmt.Sprintf(oauthGetUserDetailInfoURL, accessToken), options); err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfoDetail error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
