package context

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}

func (ctx *Context) SetQYAPIDomain(domain string) {
	ctx.QYAPIDomain = domain
}

// GetQYAPIDomain 返回企业微信API域名
//
// 默认值 https://qyapi.weixin.qq.com
func (ctx *Context) GetQYAPIDomain() string {
	return ctx.QYAPIDomain
}

func (ctx *Context) SetOpenQYAPIDomain(domain string) {
	ctx.OpenQYAPIDomain = domain
}

// GetOpenQYAPIDomain 返回企业微信开放平台域名
//
// 默认值 https://open.work.weixin.qq.com
func (ctx *Context) GetOpenQYAPIDomain() string {
	return ctx.OpenQYAPIDomain
}

func (ctx *Context) SetOpenAPIDomain(domain string) {
	ctx.OpenAPIDomain = domain
}

// GetOpenAPIDomain 返回微信开放平台域名
//
// 默认值 https://open.weixin.qq.com
func (ctx *Context) GetOpenAPIDomain() string {
	return ctx.OpenAPIDomain
}
