package agent

import (
	"github.com/silenceper/wechat/v2/util"
)

type RespGetAgent struct {
	util.CommonError
	AgentId        int    `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoUrl  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserInfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowParties struct {
		PartyId []int `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct {
		TagId []int `json:"tagid"`
	} `json:"allow_tags"`
	Close              int    `json:"close"`
	RedirectDomain     string `json:"redirect_domain"`
	ReportLocationFlag int    `json:"report_location_flag"`
	IsReportEnter      int    `json:"isreportenter"`
	HomeUrl            string `json:"home_url"`
}

type RespListAgent struct {
	util.CommonError
	AgentLists []struct {
		AgentId       int    `json:"agentid"`
		Name          string `json:"name"`
		SquareLogoUrl string `json:"square_logo_url"`
	} `json:"agentlist"`
}

type ReqAgentSet struct {
	AgentId            int    `json:"agentid"`
	ReportLocationFlag int    `json:"report_location_flag"`
	LogoMediaId        string `json:"logo_mediaid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RedirectDomain     string `json:"redirect_domain"`
	IsReportEnter      int    `json:"isreportenter"`
	HomeUrl            string `json:"home_url"`
}

type MenuButton struct {
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Key       string       `json:"key,omitempty"`
	URL       string       `json:"url,omitempty"`
	PagePath  string       `json:"pagepath,omitempty"`
	AppId     string       `json:"appid,omitempty"`
	SubButton []MenuButton `json:"sub_button,omitempty"`
}

type ReqAgentMenuCreate struct {
	Button MenuButton `json:"button"`
}

type RespAgentMenuGet struct {
	util.CommonError
	Button MenuButton `json:"button"`
}
