package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/util"
)

// getTicketURL 获取ticket的url
const (
	workGetTicketURL      = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=%s"
	workAgentGetTicketURL = "https://qyapi.weixin.qq.com/cgi-bin/ticket/get?access_token=%s&type=agent_config"
)

// WorkJsTicket 默认获取js ticket方法
type WorkJsTicket struct {
	corpID         string
	agentID        string
	cacheKeyPrefix string
	cache          cache.Cache
	// jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.Mutex
}

// NewWorkJsTicket new
func NewWorkJsTicket(corpId, agentId string, cacheKeyPrefix string, cache cache.Cache) JsTicketHandle {
	return &WorkJsTicket{
		corpID:          corpId,
		agentID:         agentId,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		jsAPITicketLock: new(sync.Mutex),
	}
}

// GetTicket 获取jsapi_ticket
func (js *WorkJsTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	jsAPITicketCacheKey := fmt.Sprintf("%s_qy_jsapi_ticket_%s", js.cacheKeyPrefix, js.corpID)
	if len(js.agentID) > 0 {
		jsAPITicketCacheKey = fmt.Sprintf("%s_qy_agent_jsapi_ticket_%s_%s", js.cacheKeyPrefix, js.corpID, js.agentID)
	}
	if val := js.cache.Get(jsAPITicketCacheKey); val != nil {
		return val.(string), nil
	}

	js.jsAPITicketLock.Lock()
	defer js.jsAPITicketLock.Unlock()

	// 双检，防止重复从微信服务器获取
	if val := js.cache.Get(jsAPITicketCacheKey); val != nil {
		return val.(string), nil
	}
	var ticket ResTicket
	if len(js.agentID) > 0 {
		ticket, err = getWorkTicketFromServer(fmt.Sprintf(workAgentGetTicketURL, accessToken))
	} else {
		ticket, err = getWorkTicketFromServer(fmt.Sprintf(workGetTicketURL, accessToken))
	}
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	err = js.cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	ticketStr = ticket.Ticket
	return
}

// getWorkTicketFromServer 从服务器中获取ticket
func getWorkTicketFromServer(url string) (ticket ResTicket, err error) {
	var response []byte
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}
	return
}
