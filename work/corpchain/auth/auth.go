package auth

import (
	context2 "context"
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/xerror"
)

const (
	transferSessionURL = "/cgi-bin/miniprogram/transfer_session?access_token=%s"
)

// Auth 上级/上游企业通过该接口转换为下级/下游企业的小程序session
type Auth struct {
	*context.Context
}

// NewAuth new auth
func NewAuth(ctx *context.Context) *Auth {
	return &Auth{ctx}
}

type ReqTransferSession struct {
	UserId     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

// RespTransferSession 登录凭证校验的返回结果
type RespTransferSession struct {
	util.CommonError
	UserId     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

// TransferSession 获取下级/下游企业小程序session
// https://developer.work.weixin.qq.com/document/path/95318
func (auth *Auth) TransferSession(userId, sessionKey string) (result RespTransferSession, err error) {
	return auth.TransferSessionContext(context2.Background(), userId, sessionKey)
}

// TransferSessionContext 上级/上游企业通过该接口转换为下级/下游企业的小程序session
func (auth *Auth) TransferSessionContext(ctx context2.Context, userId, sessionKey string) (result RespTransferSession, err error) {
	var (
		response    []byte
		accessToken string
	)
	accessToken, err = auth.GetAccessToken()
	if err != nil {
		return
	}
	req := &ReqTransferSession{
		UserId:     userId,
		SessionKey: sessionKey,
	}
	if response, err = util.PostJSONContext(ctx, auth.GetQYAPIDomain()+fmt.Sprintf(transferSessionURL, accessToken), req); err != nil {
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
