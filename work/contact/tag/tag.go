package tag

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/contact"
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	tagCreateAddr     = "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s"
	tagUpdateAddr     = "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=%s"
	tagDeleteAddr     = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=%s&tagid=%d"
	tagGetAddr        = "https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=%s&tagid=%d"
	tagAddUserAddr    = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=%s"
	tagDeleteUserAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=%s"
	tagListAddr       = "https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=%s&id=%d"
)

type Tag struct {
	ctx *context.Context
}

// NewTag 初始化企业微信标签实例
func NewTag(ctx *context.Context) (client *Tag, err error) {
	client = &Tag{
		ctx: ctx,
	}
	return client, nil
}

// Create 创建标签
// https://work.weixin.qq.com/api/doc/90000/90135/90210
func (r *Tag) Create(options ObjTag) (info RespTag, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(tagCreateAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// Update 更新标签名字
// https://work.weixin.qq.com/api/doc/90000/90135/90211
func (r *Tag) Update(options ObjTag) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(tagUpdateAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// Delete 删除标签
// https://work.weixin.qq.com/api/doc/90000/90135/90212
func (r *Tag) Delete(id int) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(tagDeleteAddr, accessToken, id))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// List 获取标签列表
// https://work.weixin.qq.com/api/doc/90000/90135/90216
func (r *Tag) List(id int) (info RespTagList, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(tagListAddr, accessToken, id))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// GetUsers 获取标签成员
// https://work.weixin.qq.com/api/doc/90000/90135/90213
func (r *Tag) GetUsers(id int) (info RespTagGet, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(tagGetAddr, accessToken, id))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// AddUsers 增加标签成员
// https://work.weixin.qq.com/api/doc/90000/90135/90214
func (r *Tag) AddUsers(options ReqTagUser) (info RespTagUser, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(tagAddUserAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// DeleteUsers 删除标签成员
// https://work.weixin.qq.com/api/doc/90000/90135/90215
func (r *Tag) DeleteUsers(options ReqTagUser) (info RespTagUser, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(tagDeleteUserAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
