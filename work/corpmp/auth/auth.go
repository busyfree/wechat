package auth

import (
	context2 "context"
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/contact/xerror"
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	code2SessionURL = "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=%s&js_code=%s&grant_type=authorization_code"
)

// Auth 登录/用户信息
type Auth struct {
	*context.Context
}

// NewAuth new auth
func NewAuth(ctx *context.Context) *Auth {
	return &Auth{ctx}
}

// RespCode2Session 登录凭证校验的返回结果
type RespCode2Session struct {
	util.CommonError
	CorpId     string `json:"corpid"`
	UserId     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

// Code2Session 登录凭证校验。
// https://work.weixin.qq.com/api/doc/90000/90136/91507
func (auth *Auth) Code2Session(jsCode string) (result RespCode2Session, err error) {
	return auth.Code2SessionContext(context2.Background(), jsCode)
}

// Code2SessionContext 登录凭证校验。
func (auth *Auth) Code2SessionContext(ctx context2.Context, jsCode string) (result RespCode2Session, err error) {
	var (
		response    []byte
		accessToken string
	)
	accessToken, err = auth.GetAccessToken()
	if err != nil {
		return
	}
	if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(code2SessionURL, accessToken, jsCode)); err != nil {
		return
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = xerror.NewSDKErr(result.ErrCode, result.ErrMsg)
		return
	}
	return
}
