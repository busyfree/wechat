/*
 * Copyright  (c) 2022 MS. All rights reserved.
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
 *
 * File:    corpchain.go
 * Created: 2022/6/7 11:40:41
 * Authors: MS<geek.snail@qq.com>
 */

package corpchain

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/corpchain/auth"
)

const (
	listAppShareInfoURL        = "/cgi-bin/corpgroup/corp/list_app_share_info?access_token=%s"
	getChainListURL            = "/cgi-bin/corpgroup/corp/get_chain_list?access_token=%s"
	unionidToExternalUseridURL = "/cgi-bin/corpgroup/unionid_to_external_userid?access_token=%s"
	ruleListIdsURL             = "/cgi-bin/corpgroup/rule/list_ids?access_token=%s"
	ruleDeleteRuleURL          = "/cgi-bin/corpgroup/rule/delete_rule?access_token=%s"
	ruleGetRuleInfoURL         = "/cgi-bin/corpgroup/rule/get_rule_info?access_token=%s"
	ruleAddRuleInfoURL         = "/cgi-bin/corpgroup/rule/add_rule?access_token=%s"
	ruleModifyRuleInfoURL      = "/cgi-bin/corpgroup/rule/modify_rule?access_token=%s"
)

// CorpChain 上下游相关API
// https://developer.work.weixin.qq.com/document/path/95308
type CorpChain struct {
	ctx *context.Context
}

// NewCorpChain 实例化上下游相关API
func NewCorpChain(ctx *context.Context) *CorpChain {
	return &CorpChain{ctx}
}

// GetAuth 返回 auth 实例
// 上级/上游企业通过该接口转换为下级/下游企业的小程序session
func (r *CorpChain) GetAuth() *auth.Auth {
	return auth.NewAuth(r.ctx)
}

// ListAppShareInfo 获取应用共享信息
// https://developer.work.weixin.qq.com/document/path/95815
func (r *CorpChain) ListAppShareInfo(req *ReqListAppShareInfo) (resp RespListAppShareInfo, err error) {
	var body []byte
	workCorpChainAccessTokenHandle, _ := r.ctx.AccessTokenHandle.(*credential.WorkCorpChainAccessToken)
	parentAccessToken, err := workCorpChainAccessTokenHandle.GetParentAccessToken()
	if err != nil {
		return
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(listAppShareInfoURL, parentAccessToken), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

// GetChainList 获取上下游列表
// https://developer.work.weixin.qq.com/document/path/95315
func (r *CorpChain) GetChainList() (resp RespGetChainList, err error) {
	var body []byte
	workCorpChainAccessTokenHandle, _ := r.ctx.AccessTokenHandle.(*credential.WorkCorpChainAccessToken)
	parentAccessToken, err := workCorpChainAccessTokenHandle.GetParentAccessToken()
	if err != nil {
		return
	}
	body, err = util.HTTPGet(r.ctx.QYAPIDomain + fmt.Sprintf(getChainListURL, parentAccessToken))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

// UnionidToExternalUserId 上下游企业应用获取微信用户的external_userid
// https://developer.work.weixin.qq.com/document/path/95342
func (r *CorpChain) UnionidToExternalUserId(unionId, openId, corpId string) (resp RespUnionidToExternalUserId, err error) {
	var body []byte
	token, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	req := &ReqUnionidToExternalUserId{
		Unionid: unionId,
		Openid:  openId,
		CorpId:  corpId,
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(unionidToExternalUseridURL, token), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func (r *CorpChain) RuleListIds(chainId string) (resp RespRuleListIds, err error) {
	var body []byte
	token, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	req := &ReqRuleListIds{
		ChainId: chainId,
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(ruleListIdsURL, token), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func (r *CorpChain) RuleDeleteRule(chainId string, ruleId int) (resp util.CommonError, err error) {
	var body []byte
	token, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	req := &ReqRuleDeleteRule{
		ChainId: chainId,
		RuleId:  ruleId,
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(ruleDeleteRuleURL, token), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func (r *CorpChain) RuleGetRuleInfo(chainId string, ruleId int) (resp RespRuleGetRuleInfo, err error) {
	var body []byte
	token, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	req := &ReqRuleDeleteRule{
		ChainId: chainId,
		RuleId:  ruleId,
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(ruleGetRuleInfoURL, token), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func (r *CorpChain) RuleAddRule(req *ReqRuleAddRule) (resp RespRuleAddRule, err error) {
	var body []byte
	token, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(ruleAddRuleInfoURL, token), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func (r *CorpChain) RuleModifyRule(req *ReqRuleModifyRule) (resp util.CommonError, err error) {
	var body []byte
	token, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	body, err = util.PostJSON(r.ctx.QYAPIDomain+fmt.Sprintf(ruleModifyRuleInfoURL, token), req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}
