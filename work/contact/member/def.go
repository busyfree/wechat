package member

import (
	"github.com/silenceper/wechat/v2/util"
)

type CommonFields struct {
	UserId         string   `json:"userid"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	Mobile         string   `json:"mobile"`
	Department     []int    `json:"department"`
	Order          []int    `json:"order"`
	Position       string   `json:"position"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	DirectLeader   []string `json:"direct_leader"`
	Telephone      string   `json:"telephone"`
	Address        string   `json:"address"`
	MainDepartment int      `json:"main_department"`
	ExtAttr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int    `json:"status,omitempty"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			MiniProgram struct {
				Appid    string `json:"appid"`
				PagePath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

// ReqMemberCreate https://work.weixin.qq.com/api/doc/90000/90135/90195
type ReqMemberCreate struct {
	Enable        int    `json:"enable"`
	AvatarMediaId string `json:"avatar_mediaid"`
	ToInvite      bool   `json:"to_invite"`
	CommonFields
}

// RespMemberRead https://work.weixin.qq.com/api/doc/90000/90135/90196
type RespMemberRead struct {
	util.CommonError
	Avatar      string `json:"avatar"`
	ThumbAvatar string `json:"thumb_avatar"`
	Status      int    `json:"status"`
	QrCode      string `json:"qr_code"`
	CommonFields
}

// ReqMemberUpdate https://work.weixin.qq.com/api/doc/90000/90135/90197
type ReqMemberUpdate struct {
	Enable        int    `json:"enable"`
	AvatarMediaId string `json:"avatar_mediaid"`
	CommonFields
}

// ReqMemberBatchDelete https://work.weixin.qq.com/api/doc/90000/90135/90199
type ReqMemberBatchDelete struct {
	UserIdList []string `json:"useridlist"`
}

// RespMemberDepartmentMembers https://work.weixin.qq.com/api/doc/90000/90135/90200
type RespMemberDepartmentMembers struct {
	util.CommonError
	UserList []struct {
		Userid     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserId string `json:"open_userid"`
	} `json:"userlist"`
}

// RespMemberDepartmentMemberDetail https://work.weixin.qq.com/api/doc/90000/90135/90201
type RespMemberDepartmentMemberDetail struct {
	util.CommonError
	UserList []struct {
		Avatar      string `json:"avatar"`
		ThumbAvatar string `json:"thumb_avatar"`
		Status      int    `json:"status"`
		HideMobile  int    `json:"hide_mobile"`
		EnglishName string `json:"english_name"`
		OpenUserId  string `json:"open_userid"`
		QRCode      string `json:"qr_code"`
		CommonFields
	} `json:"userlist"`
}

// ReqMemberConvertToOpenId https://work.weixin.qq.com/api/doc/90000/90135/90202
type ReqMemberConvertToOpenId struct {
	UserId string `json:"userid"`
}

type RespMemberConvertToOpenId struct {
	util.CommonError
	OpenId string `json:"openid"`
}

// ReqMemberBatchInvite https://work.weixin.qq.com/api/doc/90000/90135/90975
type ReqMemberBatchInvite struct {
	User  []string      `json:"user"`
	Party []interface{} `json:"party"`
	Tag   []interface{} `json:"tag"`
}
type RespMemberBatchInvite struct {
	util.CommonError
	InvalidUser  []string      `json:"invaliduser"`
	InvalidParty []interface{} `json:"invalidparty"`
	InvalidTag   []interface{} `json:"invalidtag"`
}

// RespMemberGetJoinQRCode https://work.weixin.qq.com/api/doc/90000/90135/91714
type RespMemberGetJoinQRCode struct {
	util.CommonError
	JoinQRCode string `json:"join_qrcode"`
}

// ReqMemberGetActiveStat https://work.weixin.qq.com/api/doc/90000/90135/92714
type ReqMemberGetActiveStat struct {
	Date string `json:"date"`
}
type RespMemberGetActiveStat struct {
	util.CommonError
	ActiveCnt int `json:"active_cnt"`
}

// AsyncBatchSyncUserReq https://work.weixin.qq.com/api/doc/90000/90135/90980
type AsyncBatchSyncUserReq struct {
	MediaId  string `json:"media_id"`
	ToInvite bool   `json:"to_invite"`
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

type BatchAsyncUpdateOrReplaceUser struct {
	UserId string `json:"userid"`
	util.CommonError
}

type BatchGetAsyncJobResultResp struct {
	util.CommonError
	Status     int                             `json:"status"`
	Type       string                          `json:"type"`
	Total      int                             `json:"total"`
	Percentage int                             `json:"percentage"`
	Result     []BatchAsyncUpdateOrReplaceUser `json:"result"`
}

type ReqAsyncExportUser struct {
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
