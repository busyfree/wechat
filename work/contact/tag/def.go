package tag

import (
	"github.com/silenceper/wechat/v2/util"
)

type ObjTag struct {
	TagName string `json:"tagname"`
	TagId   int    `json:"tagid,omitempty"`
}

type RespTag struct {
	util.CommonError
	TagId int `json:"tagid"`
}

type RespTagGet struct {
	util.CommonError
	TagName  string `json:"tagname"`
	UserList []struct {
		UserId string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
	PartyList []int `json:"partylist"`
}

type ReqTagUser struct {
	TagId     int      `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int    `json:"partylist"`
}

type RespTagUser struct {
	util.CommonError
	InvalidList  string `json:"invalidlist"`
	InvalidParty []int  `json:"invalidparty"`
}

type RespTagList struct {
	util.CommonError
	TagList []struct {
		TagId   int    `json:"tagid"`
		TagName string `json:"tagname"`
	} `json:"taglist"`
}
