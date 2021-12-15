package contact

import (
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/contact/department"
	"github.com/silenceper/wechat/v2/work/contact/linkedcorp"
	"github.com/silenceper/wechat/v2/work/contact/member"
	"github.com/silenceper/wechat/v2/work/contact/notify"
	"github.com/silenceper/wechat/v2/work/contact/tag"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/xerror"
)

// Contact 企业微信通讯录相关API
type Contact struct {
	corpID         string // 企业ID：企业开通的每个微信客服，都对应唯一的企业ID，企业可在微信客服管理后台的企业信息处查看
	secret         string // Secret是微信客服用于校验开发者身份的访问密钥，企业成功注册微信客服后，可在「微信客服管理后台-开发配置」处获取
	token          string // 用于生成签名校验回调请求的合法性
	encodingAESKey string // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	cache          cache.Cache
	ctx            *context.Context
}

// NewContact 实例化企业微信通讯录API
func NewContact(cfg *config.Config) (client *Contact, err error) {
	if cfg.Cache == nil {
		return nil, xerror.NewSDKErr(50001)
	}
	// 初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Contact{
		corpID:         cfg.CorpID,
		secret:         cfg.CorpSecret,
		token:          cfg.Token,
		encodingAESKey: cfg.EncodingAESKey,
		cache:          cfg.Cache,
		ctx:            ctx,
	}

	return client, nil
}

// SetAccessTokenHandle 自定义access_token获取方式
func (r *Contact) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	r.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (r *Contact) GetContext() *context.Context {
	return r.ctx
}

// GetDepartment 返回 department 实例
func (r *Contact) GetDepartment() *department.Department {
	return department.NewDepartment(r.ctx)
}

// GetLinkedCorp 返回 linked corp 实例
func (r *Contact) GetLinkedCorp() *linkedcorp.LinkedCorp {
	return linkedcorp.NewLinkedCorp(r.ctx)
}

// GetMember 返回 member 实例
func (r *Contact) GetMember() *member.Member {
	return member.NewMember(r.ctx)
}

// GetTag 返回 tag 实例
func (r *Contact) GetTag() *tag.Tag {
	return tag.NewTag(r.ctx)
}

// GetNotify 返回 notify 实例
func (r *Contact) GetNotify() (*notify.Notify, error) {
	return notify.NewNotify(r.ctx.Config)
}
