package member

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/contact"
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	userCreateAddr                    = "https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s"
	userReadAddr                      = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
	userUpdateAddr                    = "https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=%s"
	userDeleteAddr                    = "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%s"
	userBatchDeleteAddr               = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=%s"
	userSimpleListAddr                = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%s&fetch_child=%s"
	userListAddr                      = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%s&fetch_child=%s"
	userConvertToOpenIdAdd            = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=%s"
	userAuthSuccessAddr               = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=%s&userid=%s"
	userBatchInviteAddr               = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite?access_token=%s"
	userGetCorpJoinQRCodeAddr         = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode?access_token=%s&size_type=%d"
	userGetActiveStatAddr             = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat?access_token=%s"
	userAsyncBatchSyncUpdateUserAddr  = "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=%s"
	userAsyncBatchSyncReplaceUserAddr = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=%s"
	getAsyncJobResultAddr             = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=%s&jobid=%s"
	getAsyncExportJobResultAddr       = "https://qyapi.weixin.qq.com/cgi-bin/export/get_result?access_token=%s&jobid=%s"
	userAsyncExportSimpleUserAddr     = "https://qyapi.weixin.qq.com/cgi-bin/export/simple_user?access_token=%s"
	userAsyncExportUserAddr           = "https://qyapi.weixin.qq.com/cgi-bin/export/user?access_token=%s"
	userAsyncExportTagUsersAddr       = "https://qyapi.weixin.qq.com/cgi-bin/export/taguser?access_token=%s"
)

type Member struct {
	ctx *context.Context
}

// NewMember 初始化企业微信成员实例
func NewMember(ctx *context.Context) (client *Member, err error) {
	client = &Member{
		ctx: ctx,
	}
	return client, nil
}

// Create 创建成员
// https://work.weixin.qq.com/api/doc/90000/90135/90195
func (r *Member) Create(options ReqMemberCreate) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userCreateAddr, accessToken), options)
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

// Read 读取成员
// https://work.weixin.qq.com/api/doc/90000/90135/90196
func (r *Member) Read(userId string) (info RespMemberRead, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userReadAddr, accessToken, userId))
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

// Update 更新成员
// https://work.weixin.qq.com/api/doc/90000/90135/90197
func (r *Member) Update(options ReqMemberUpdate) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userUpdateAddr, accessToken), options)
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

// Delete 删除成员
// https://work.weixin.qq.com/api/doc/90000/90135/90198
func (r *Member) Delete(userId string) (info RespMemberRead, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userDeleteAddr, accessToken, userId))
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

// DeleteBatchUserIds 批量删除成员
// https://work.weixin.qq.com/api/doc/90000/90135/90199
func (r *Member) DeleteBatchUserIds(options ReqMemberBatchDelete) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	if len(options.UserIdList) > 200 {
		err = errors.New("userid list length large than 200")
		return
	}
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userBatchDeleteAddr, accessToken), options)
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

// GetMemberSimpleList 获取部门成员
// https://work.weixin.qq.com/api/doc/90000/90135/90200
func (r *Member) GetMemberSimpleList(departmentId, fetchChild string) (info RespMemberDepartmentMembers, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userSimpleListAddr, accessToken, departmentId, fetchChild))
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

// GetMemberList 获取部门成员详情
// https://work.weixin.qq.com/api/doc/90000/90135/90201
func (r *Member) GetMemberList(departmentId, fetchChild string) (info RespMemberDepartmentMemberDetail, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userListAddr, accessToken, departmentId, fetchChild))
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

// ConvertToOpenId userid与openid互换
// https://work.weixin.qq.com/api/doc/90000/90135/90202
func (r *Member) ConvertToOpenId(options ReqMemberConvertToOpenId) (info RespMemberConvertToOpenId, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userConvertToOpenIdAdd, accessToken), options)
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

// AuthSuccess 二次验证 如果成员是首次加入企业，企业获取到userid，并验证了成员信息后，调用如下接口即可让成员成功加入企业。
// https://work.weixin.qq.com/api/doc/90000/90135/90203
func (r *Member) AuthSuccess(userId string) (info RespMemberDepartmentMemberDetail, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userAuthSuccessAddr, accessToken, userId))
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

// BatchInvite 邀请成员
// https://work.weixin.qq.com/api/doc/90000/90135/90975
func (r *Member) BatchInvite(options ReqMemberBatchInvite) (info RespMemberBatchInvite, err error) {
	var (
		accessToken string
		data        []byte
	)
	if len(options.User) > 1000 {
		err = errors.New("user ids large than 1000")
		return
	}
	if len(options.Party) > 100 || len(options.Tag) > 100 {
		err = errors.New("user party or tag large than 100")
		return
	}
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userBatchInviteAddr, accessToken), options)
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

// GetJoinQRCode 获取加入企业二维码
// https://work.weixin.qq.com/api/doc/90000/90135/91714
func (r *Member) GetJoinQRCode(sizeType int) (info RespMemberGetJoinQRCode, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userGetCorpJoinQRCodeAddr, accessToken, sizeType))
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

// GetActiveStat 获取企业活跃成员数
// https://work.weixin.qq.com/api/doc/90000/90135/92714
func (r *Member) GetActiveStat(options ReqMemberGetActiveStat) (info RespMemberGetActiveStat, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userGetActiveStatAddr, accessToken), options)
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

// AsyncBatchSyncUpdateUser 增量更新成员
// https://work.weixin.qq.com/api/doc/90000/90135/90980
func (r *Member) AsyncBatchSyncUpdateUser(options AsyncBatchSyncUserReq) (info AsyncJobResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userAsyncBatchSyncUpdateUserAddr, accessToken), options)
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

// AsyncBatchSyncReplaceUser 全量覆盖成员
// https://work.weixin.qq.com/api/doc/90000/90135/90981
func (r *Member) AsyncBatchSyncReplaceUser(options AsyncBatchSyncUserReq) (info AsyncJobResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userAsyncBatchSyncReplaceUserAddr, accessToken), options)
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

// BatchGetAsyncJobResult 获取异步任务结果
// https://work.weixin.qq.com/api/doc/90000/90135/90983
func (r *Member) BatchGetAsyncJobResult(jobId string) (info BatchGetAsyncJobResultResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getAsyncJobResultAddr, accessToken, jobId))
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

// AsyncExportSimpleUser 导出成员
// https://work.weixin.qq.com/api/doc/90000/90135/94849
func (r *Member) AsyncExportSimpleUser(options ReqAsyncExportUser) (info AsyncJobResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	if options.BlockSize > 100_0000 {
		err = errors.New("block_size too large")
		return
	}
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userAsyncExportSimpleUserAddr, accessToken), options)
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

// AsyncExportUser 导出成员详情
// https://work.weixin.qq.com/api/doc/90000/90135/94851
func (r *Member) AsyncExportUser(options ReqAsyncExportUser) (info AsyncJobResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	if options.BlockSize > 100_0000 {
		err = errors.New("block_size too large")
		return
	}
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userAsyncExportUserAddr, accessToken), options)
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

// AsyncExportTagUsers 导出标签成员
// https://work.weixin.qq.com/api/doc/90000/90135/94853
func (r *Member) AsyncExportTagUsers(options ReqAsyncExportUser) (info AsyncJobResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	if options.BlockSize > 100_0000 {
		err = errors.New("block_size too large")
		return
	}
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userAsyncExportTagUsersAddr, accessToken), options)
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

// GetAsyncExportJobResult 获取导出结果 获取任务结果的调用身份需要与提交任务的一致
// https://work.weixin.qq.com/api/doc/90000/90135/94854
func (r *Member) GetAsyncExportJobResult(jobId string) (info GetAsyncExportJobResultResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getAsyncExportJobResultAddr, accessToken, jobId))
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
