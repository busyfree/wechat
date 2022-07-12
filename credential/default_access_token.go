package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// AccessTokenURL 获取access_token的接口
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// AccessTokenURL 企业微信获取access_token的接口
	workAccessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	// 获取应用可见范围内下级/下游企业的access_token，该access_token可用于调用下级/下游企业通讯录的只读接口。
	workCorpChainAccessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/gettoken?access_token=%s"
	// CacheKeyOfficialAccountPrefix 微信公众号cache key前缀
	CacheKeyOfficialAccountPrefix = "gowechat_officialaccount_"
	// CacheKeyMiniProgramPrefix 小程序cache key前缀
	CacheKeyMiniProgramPrefix = "gowechat_miniprogram_"
	// CacheKeyWorkPrefix 企业微信cache key前缀
	CacheKeyWorkPrefix = "gowechat_work_"
)

// DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID           string
	appSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is ineed")
	}
	return &DefaultAccessToken{
		appID:           appID,
		appSecret:       appSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetAccessToken 获取access_token,先从cache中获取，没有则从服务端获取
func (ak *DefaultAccessToken) GetAccessToken() (accessToken string, err error) {
	// 先从cache中取
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	if val := ak.cache.Get(accessTokenCacheKey); val != nil {
		return val.(string), nil
	}

	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	// 双检，防止重复从微信服务器获取
	if val := ak.cache.Get(accessTokenCacheKey); val != nil {
		return val.(string), nil
	}

	// cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(fmt.Sprintf(accessTokenURL, ak.appID, ak.appSecret))
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

// WorkAccessToken 企业微信AccessToken 获取
type WorkAccessToken struct {
	CorpID          string
	CorpSecret      string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewWorkAccessToken new WorkAccessToken
func NewWorkAccessToken(corpID, corpSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache the not exist")
	}
	return &WorkAccessToken{
		CorpID:          corpID,
		CorpSecret:      corpSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// GetAccessToken 企业微信获取access_token,先从cache中获取，没有则从服务端获取
func (ak *WorkAccessToken) GetAccessToken() (accessToken string, err error) {
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()
	corpSecretMd5Key, _ := util.CalculateSign(ak.CorpSecret, "MD5", "")
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s_%s", ak.cacheKeyPrefix, ak.CorpID, corpSecretMd5Key)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	// cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(fmt.Sprintf(workAccessTokenURL, ak.CorpID, ak.CorpSecret))
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

// WorkCorpChainAccessToken 企业微信AccessToken 获取
type WorkCorpChainAccessToken struct {
	CorpID                string
	AgentID               int
	BusinessType          int
	cacheKeyPrefix        string
	cache                 cache.Cache
	accessTokenLock       *sync.Mutex
	parentCorpAccessToken *WorkAccessToken
}

// NewWorkCorpChainAccessToken new WorkCorpChainAccessToken
func NewWorkCorpChainAccessToken(parentCorpAccessTokenHandle AccessTokenHandle, chainCorpId string, agentId int, cacheKeyPrefix string, bizType int, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache the not exist")
	}
	parentWorkAccessToken, ok := parentCorpAccessTokenHandle.(*WorkAccessToken)
	if !ok {
		panic("missing parentCorpAccessTokenHandle")
	}
	return &WorkCorpChainAccessToken{
		parentCorpAccessToken: parentWorkAccessToken,
		CorpID:                chainCorpId,
		AgentID:               agentId,
		BusinessType:          bizType,
		cache:                 cache,
		cacheKeyPrefix:        cacheKeyPrefix,
		accessTokenLock:       new(sync.Mutex),
	}
}

// GetAccessToken 获取下级/下游企业的access_token,先从cache中获取，没有则从服务端获取
func (ak *WorkCorpChainAccessToken) GetAccessToken() (accessToken string, err error) {
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()
	accessTokenCacheKey := fmt.Sprintf("%s_chain_access_token_%s_%d", ak.cacheKeyPrefix, ak.CorpID, ak.AgentID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}
	// cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = ak.getWorkCorpChainTokenFromServer()
	if err != nil {
		return
	}
	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

func (ak *WorkCorpChainAccessToken) GetParentAccessToken() (accessToken string, err error) {
	accessToken, err = ak.parentCorpAccessToken.GetAccessToken()
	if err != nil {
		return
	}
	return
}

type ReqWorkCorpChainToken struct {
	CorpId  string `json:"corpid"`
	BizType int    `json:"business_type"`
	AgentId int    `json:"agentid"`
}

func (ak *WorkCorpChainAccessToken) getWorkCorpChainTokenFromServer() (resAccessToken ResAccessToken, err error) {
	var body []byte
	req := &ReqWorkCorpChainToken{
		CorpId:  ak.CorpID,
		BizType: ak.BusinessType,
		AgentId: ak.AgentID,
	}
	parentAccessToken, err := ak.parentCorpAccessToken.GetAccessToken()
	if err != nil {
		return
	}
	body, err = util.PostJSON(fmt.Sprintf(workCorpChainAccessTokenURL, parentAccessToken), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}

// GetTokenFromServer 强制从微信服务器获取token
func GetTokenFromServer(url string) (resAccessToken ResAccessToken, err error) {
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}
