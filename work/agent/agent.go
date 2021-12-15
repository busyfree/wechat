package agent

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/xerror"
)

const (
	agentGetAddr        = "https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=%s&agentid=%d"
	agentListAddr       = "https://qyapi.weixin.qq.com/cgi-bin/agent/list?access_token=%s"
	agentSetAddr        = "https://qyapi.weixin.qq.com/cgi-bin/agent/set?access_token=%s"
	agentMenuCreateAddr = "https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=%s&agentid=%d"
	agentMenuGetAddr    = "https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=%s&agentid=%d"
	agentMenuDeleteAddr = "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=%s&agentid=%d"
)

// Agent 应用管理API
type Agent struct {
	ctx *context.Context
}

// NewAgent 实例化应用管理API
func NewAgent(ctx *context.Context) *Agent {
	return &Agent{ctx}
}

// Get 获取指定的应用详情
// https://work.weixin.qq.com/api/doc/90000/90135/90227
func (r *Agent) Get(agentId int32) (info RespGetAgent, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(agentGetAddr, accessToken, agentId))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, xerror.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// List 获取access_token对应的应用列表
// https://work.weixin.qq.com/api/doc/90000/90135/90227
func (r *Agent) List() (info RespListAgent, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(agentListAddr, accessToken))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, xerror.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// Set 设置应用
// https://work.weixin.qq.com/api/doc/90000/90135/90228
func (r *Agent) Set(options ReqAgentSet) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(agentSetAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, xerror.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// MenuCreate 创建菜单
// https://work.weixin.qq.com/api/doc/90000/90135/90231
func (r *Agent) MenuCreate(agentId int32, options ReqAgentMenuCreate) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	if len(options.Button) > 3 {
		err = errors.New("一级菜单数组，个数应为1~3个")
		return
	}
	for _, b := range options.Button {
		if len(b.SubButton) > 5 {
			err = errors.New("二级菜单数组，个数应为1~5个")
			return
		}
	}
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(agentMenuCreateAddr, accessToken, agentId), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, xerror.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// MenuGet 获取菜单
// https://work.weixin.qq.com/api/doc/90000/90135/90232
func (r *Agent) MenuGet(agentId int32) (info RespAgentMenuGet, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(agentMenuGetAddr, accessToken, agentId))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, xerror.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// MenuDelete 删除菜单
// https://work.weixin.qq.com/api/doc/90000/90135/90233
func (r *Agent) MenuDelete(agentId int32) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(agentMenuDeleteAddr, accessToken, agentId))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, xerror.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
