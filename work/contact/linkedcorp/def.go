package linkedcorp

import (
	"github.com/silenceper/wechat/v2/util"
)

type AgentGetPermListResp struct {
	util.CommonError
	UserIds       []string `json:"userids"`
	DepartmentIds []string `json:"department_ids"`
}

type GetUserDetailReq struct {
	UserId string `json:"userid"`
}

type GetUserDetailResp struct {
	util.CommonError
	UserInfo struct {
		UserId     string   `json:"userid"`
		Name       string   `json:"name"`
		Department []string `json:"department"`
		Mobile     string   `json:"mobile"`
		Telephone  string   `json:"telephone"`
		Email      string   `json:"email"`
		Position   string   `json:"position"`
		CorpId     string   `json:"corpid"`
		ExtAttr    struct {
			Attrs []struct {
				Name  string `json:"name"`
				Value string `json:"value,omitempty"`
				Type  int    `json:"type"`
				Text  struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					Url   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
			} `json:"attrs"`
		} `json:"extattr"`
	} `json:"user_info"`
}

type GetUserListReq struct {
	DepartmentId string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}

type GetUserSimpleListResp struct {
	util.CommonError
	UserList []struct {
		Userid     string   `json:"userid"`
		Name       string   `json:"name"`
		Department []string `json:"department"`
		CorpId     string   `json:"corpid"`
	} `json:"userlist"`
}

type GetUserListResp struct {
	util.CommonError
	UserList []struct {
		UserId     string   `json:"userid"`
		Name       string   `json:"name"`
		Department []string `json:"department"`
		Mobile     string   `json:"mobile"`
		Telephone  string   `json:"telephone"`
		Email      string   `json:"email"`
		Position   string   `json:"position"`
		CorpId     string   `json:"corpid"`
		ExtAttr    struct {
			Attrs []struct {
				Name  string `json:"name"`
				Value string `json:"value,omitempty"`
				Type  int    `json:"type"`
				Text  struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					Url   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
			} `json:"attrs"`
		} `json:"extattr"`
	} `json:"userlist"`
}

type GetDepartmentListReq struct {
	DepartmentId string `json:"department_id"`
}

type GetDepartmentListResp struct {
	util.CommonError
	DepartmentList []struct {
		DepartmentId   string `json:"department_id"`
		DepartmentName string `json:"department_name"`
		ParentId       string `json:"parentid"`
		Order          int    `json:"order"`
	} `json:"department_list"`
}
