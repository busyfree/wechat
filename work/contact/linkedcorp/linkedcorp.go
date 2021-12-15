package linkedcorp

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/xerror"
)

const (
	linkedCorpAgentAddr          = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list?access_token=%s"
	linkedCorpUserDetailAddr     = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get?access_token=%s"
	linkedCorpUserSimpleListAddr = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/simplelist?access_token=%s"
	linkedCorpUserListAddr       = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/list?access_token=%s"
	linkedCorpDepartmentListAddr = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list?access_token=%s"
)

type LinkedCorp struct {
	ctx *context.Context
}

// NewLinkedCorp 初始化企业微信互联企业实例
func NewLinkedCorp(ctx *context.Context) (client *LinkedCorp) {
	client = &LinkedCorp{
		ctx: ctx,
	}
	return client
}

// GetPermList 获取应用的可见范围
// https://work.weixin.qq.com/api/doc/90000/90135/93172
func (r *LinkedCorp) GetPermList() (info AgentGetPermListResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(linkedCorpAgentAddr, accessToken), nil)
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

// GetLinkedCorpUserDetail 获取互联企业成员详细信息
// https://work.weixin.qq.com/api/doc/90000/90135/93171
func (r *LinkedCorp) GetLinkedCorpUserDetail(options GetUserDetailReq) (info GetUserDetailResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(linkedCorpUserDetailAddr, accessToken), options)
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

// GetLinkedCorpDepartmentUsers 获取互联企业部门成员
// https://work.weixin.qq.com/api/doc/90000/90135/93168
func (r *LinkedCorp) GetLinkedCorpDepartmentUsers(options GetUserListReq) (info GetUserSimpleListResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(linkedCorpUserSimpleListAddr, accessToken), options)
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

// GetLinkedCorpDepartmentUserDetail 获取互联企业部门成员详情
// https://work.weixin.qq.com/api/doc/90000/90135/93169
func (r *LinkedCorp) GetLinkedCorpDepartmentUserDetail(options GetUserListReq) (info GetUserListResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(linkedCorpUserListAddr, accessToken), options)
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

// GetDepartmentList 获取互联企业部门列表
// https://work.weixin.qq.com/api/doc/90000/90135/93170
func (r *LinkedCorp) GetDepartmentList(options GetDepartmentListReq) (info GetDepartmentListResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(linkedCorpDepartmentListAddr, accessToken), options)
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
