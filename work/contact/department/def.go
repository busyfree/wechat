package department

import (
	"github.com/silenceper/wechat/v2/util"
)

type CommonField struct {
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	ParentId         int      `json:"parentid"`
	Order            int      `json:"order"`
	Id               int      `json:"id"`
	DepartmentLeader []string `json:"department_leader,omitempty"`
}

// ReqDepartmentCreate https://work.weixin.qq.com/api/doc/90000/90135/90205
type ReqDepartmentCreate struct {
	CommonField
}

type RespDepartmentCreate struct {
	util.CommonError
	Id int `json:"id"`
}

type ReqDepartmentUpdate struct {
	CommonField
}

type RespDepartmentList struct {
	util.CommonError
	Department []struct {
		CommonField
	} `json:"department"`
}

type RespDepartmentGet struct {
	util.CommonError
	Department CommonField `json:"department"`
}

type BatchAsyncReplacePartyReq struct {
	MediaId  string `json:"media_id"`
	Callback struct {
		Url            string `json:"url"`
		Token          string `json:"token"`
		EncodingAESKey string `json:"encodingaeskey"`
	} `json:"callback"`
}

type AsyncJobResp struct {
	util.CommonError
	JobId string `json:"jobid"`
}

type BatchAsyncReplaceDepartment struct {
	Action  int `json:"action"`
	PartyId int `json:"partyid"`
	util.CommonError
}

type BatchGetAsyncJobResultResp struct {
	util.CommonError
	Status     int                           `json:"status"`
	Type       string                        `json:"type"`
	Total      int                           `json:"total"`
	Percentage int                           `json:"percentage"`
	Result     []BatchAsyncReplaceDepartment `json:"result"`
}

type ReqAsyncExportDepartment struct {
	TagId          int    `json:"tagid,omitempty"`
	EncodingAESKey string `json:"encoding_aeskey"`
	BlockSize      int    `json:"block_size"`
}

type GetAsyncExportJobResultResp struct {
	util.CommonError
	Status   int `json:"status"`
	DataList []struct {
		Url  string      `json:"url"`
		Size interface{} `json:"size"`
		Md5  string      `json:"md5"`
	} `json:"data_list"`
}
