package corpmp

import (
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/corpmp/auth"
)

// CorpMP 微信小程序相关API
type CorpMP struct {
	ctx *context.Context
}

// NewCorpMP 实例化小程序API
func NewCorpMP(ctx *context.Context) *CorpMP {
	return &CorpMP{ctx}
}

// GetAuth 返回 auth 实例
func (r *CorpMP) GetAuth() *auth.Auth {
	return auth.NewAuth(r.ctx)
}
