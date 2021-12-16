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

func (ctx *Context) SetAPIDomain(domain string) {
	ctx.APIDomain = domain
}

func (ctx *Context) GetAPIDomain() string {
	return ctx.APIDomain
}
