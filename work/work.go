package work

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/agent"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/contact"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/corpmp"
	"github.com/silenceper/wechat/v2/work/kf"
	"github.com/silenceper/wechat/v2/work/msgaudit"
	"github.com/silenceper/wechat/v2/work/oauth"
)

const (
	workDefaultApiDomain     = "https://qyapi.weixin.qq.com"
	openWorkDefaultApiDomain = "https://open.work.weixin.qq.com"
	openDefaultApiDomain     = "https://open.weixin.qq.com"
)

// Work 企业微信
type Work struct {
	ctx *context.Context
}

// NewWork init work
func NewWork(cfg *config.Config) *Work {
	if cfg.Cache == nil {
		panic("cache 未设置")
	}
	if len(cfg.QYAPIDomain) == 0 {
		cfg.QYAPIDomain = workDefaultApiDomain
	}
	if len(cfg.OpenQYAPIDomain) == 0 {
		cfg.OpenQYAPIDomain = openWorkDefaultApiDomain
	}
	if len(cfg.OpenAPIDomain) == 0 {
		cfg.OpenAPIDomain = openDefaultApiDomain
	}
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Work{ctx: ctx}
}

// GetContext get Context
func (wk *Work) GetContext() *context.Context {
	return wk.ctx
}

// GetOauth get oauth
func (wk *Work) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wk.ctx)
}

// GetMsgAudit get msgAudit
func (wk *Work) GetMsgAudit() (*msgaudit.Client, error) {
	return msgaudit.NewClient(wk.ctx)
}

// GetKF get kf
func (wk *Work) GetKF() *kf.Client {
	return kf.NewClient(wk.ctx)
}

// GetContact get contact
func (wk *Work) GetContact() *contact.Contact {
	return contact.NewContact(wk.ctx)
}

// GetAgent get agent
func (wk *Work) GetAgent() *agent.Agent {
	return agent.NewAgent(wk.ctx)
}

// GetCorpMP get corp mp
func (wk *Work) GetCorpMP() *corpmp.CorpMP {
	return corpmp.NewCorpMP(wk.ctx)
}
