package department

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/xerror"
)

const (
	departmentCreateAddr                = "/cgi-bin/department/create?access_token=%s"
	departmentUpdateAddr                = "/cgi-bin/department/update?access_token=%s"
	departmentDeleteAddr                = "/cgi-bin/department/delete?access_token=%s&id=%d"
	departmentGetAddr                   = "/cgi-bin/department/get?access_token=%s&id=%d"
	departmentListAddr                  = "/cgi-bin/department/list?access_token=%s&id=%d"
	departmentAsyncReplacePartyListAddr = "/cgi-bin/batch/replaceparty?access_token=%s"
	getAsyncJobResultAddr               = "/cgi-bin/batch/getresult?access_token=%s&jobid=%s"
	departmentAsyncExportAddr           = "/cgi-bin/export/department?access_token=%s"
	getAsyncExportJobResultAddr         = "/cgi-bin/export/get_result?access_token=%s&jobid=%s"
)

type Department struct {
	ctx *context.Context
}

// NewDepartment 初始化企业微信部门实例
func NewDepartment(ctx *context.Context) (client *Department) {
	client = &Department{
		ctx: ctx,
	}
	return client
}

// Create 创建部门
// https://work.weixin.qq.com/api/doc/90000/90135/90205
func (r *Department) Create(options ReqDepartmentCreate) (info RespDepartmentCreate, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(r.ctx.GetQYAPIDomain()+fmt.Sprintf(departmentCreateAddr, accessToken), options)
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

// Update 更新部门
// https://work.weixin.qq.com/api/doc/90000/90135/90206
func (r *Department) Update(options ReqDepartmentUpdate) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(r.ctx.GetQYAPIDomain()+fmt.Sprintf(departmentUpdateAddr, accessToken), options)
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

// Delete 删除部门
// https://work.weixin.qq.com/api/doc/90000/90135/90207
func (r *Department) Delete(id int) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(r.ctx.GetQYAPIDomain() + fmt.Sprintf(departmentDeleteAddr, accessToken, id))
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

// List 获取部门列表
// https://work.weixin.qq.com/api/doc/90000/90135/90208
func (r *Department) List(id int) (info RespDepartmentList, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(r.ctx.GetQYAPIDomain() + fmt.Sprintf(departmentListAddr, accessToken, id))
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

// Get 获取单个部门详情
// https://developer.work.weixin.qq.com/document/path/95351
func (r *Department) Get(id int) (info RespDepartmentGet, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(r.ctx.GetQYAPIDomain() + fmt.Sprintf(departmentGetAddr, accessToken, id))
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

// AsyncBatchReplaceDepartment 全量覆盖部门
// https://work.weixin.qq.com/api/doc/90000/90135/90982
func (r *Department) AsyncBatchReplaceDepartment(options BatchAsyncReplacePartyReq) (info AsyncJobResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(r.ctx.GetQYAPIDomain()+fmt.Sprintf(departmentAsyncReplacePartyListAddr, accessToken), options)
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

// BatchGetAsyncJobResult 获取异步任务结果
// https://work.weixin.qq.com/api/doc/90000/90135/90983
func (r *Department) BatchGetAsyncJobResult(jobId string) (info BatchGetAsyncJobResultResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(r.ctx.GetQYAPIDomain() + fmt.Sprintf(getAsyncJobResultAddr, accessToken, jobId))
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

// AsyncExportDepartment 导出部门
// https://work.weixin.qq.com/api/doc/90000/90135/94852
func (r *Department) AsyncExportDepartment(options ReqAsyncExportDepartment) (info AsyncJobResp, err error) {
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
	data, err = util.PostJSON(r.ctx.GetQYAPIDomain()+fmt.Sprintf(departmentAsyncExportAddr, accessToken), options)
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

// GetAsyncExportJobResult 获取导出结果 获取任务结果的调用身份需要与提交任务的一致
// https://work.weixin.qq.com/api/doc/90000/90135/94854
func (r *Department) GetAsyncExportJobResult(jobId string) (info GetAsyncExportJobResultResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(r.ctx.GetQYAPIDomain() + fmt.Sprintf(getAsyncExportJobResultAddr, accessToken, jobId))
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
