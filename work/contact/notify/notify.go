package notify

import (
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/contact"
	"github.com/silenceper/wechat/v2/work/context"
)

// Notify 微信客服实例
type Notify struct {
	corpID         string // 企业ID：企业开通的每个微信客服，都对应唯一的企业ID，企业可在微信客服管理后台的企业信息处查看
	secret         string // Secret是微信客服用于校验开发者身份的访问密钥，企业成功注册微信客服后，可在「微信客服管理后台-开发配置」处获取
	token          string // 用于生成签名校验回调请求的合法性
	encodingAESKey string // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	cache          cache.Cache
	ctx            *context.Context
}

// NewNotify 初始化通讯录回调实例
func NewNotify(cfg *config.Config) (client *Notify, err error) {
	if cfg.Cache == nil {
		return nil, contact.NewSDKErr(50001)
	}
	// 初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Notify{
		corpID:         cfg.CorpID,
		secret:         cfg.CorpSecret,
		token:          cfg.Token,
		encodingAESKey: cfg.EncodingAESKey,
		cache:          cfg.Cache,
		ctx:            ctx,
	}

	return client, nil
}
