package corpchain

import (
	"github.com/silenceper/wechat/v2/util"
)

type ReqListAppShareInfo struct {
	AgentId *int    `json:"agentid"`
	BizType *int    `json:"business_type"`
	CorpId  *string `json:"corpid"`
	Limit   *int    `json:"limit"`
	Cursor  *string `json:"cursor"`
}

type RespListAppShareInfo struct {
	util.CommonError
	Ending   int `json:"ending"`
	CorpList []struct {
		CorpId   string `json:"corpid"`
		CorpName string `json:"corp_name"`
		AgentId  int    `json:"agentid"`
	} `json:"corp_list"`
	NextCursor string `json:"next_cursor"`
}

type RespGetChainList struct {
	util.CommonError
	Chains []struct {
		ChainId   string `json:"chain_id"`
		ChainName string `json:"chain_name"`
	} `json:"chains"`
}

type ReqUnionidToExternalUserId struct {
	Unionid string `json:"unionid"`
	Openid  string `json:"openid"`
	CorpId  string `json:"corpid"`
}

type RespUnionidToExternalUserId struct {
	util.CommonError
	ExternalUseridInfo []struct {
		CorpId         string `json:"corpid"`
		ExternalUserid string `json:"external_userid"`
	} `json:"external_userid_info"`
}

type ReqRuleListIds struct {
	ChainId string `json:"chain_id"`
}

type RespRuleListIds struct {
	util.CommonError
	RuleIds []int `json:"rule_ids"`
}

type ReqRuleDeleteRule struct {
	ChainId string `json:"chain_id"`
	RuleId  int    `json:"rule_id"`
}

type ReqRuleGetRuleInfo struct {
	ChainId string `json:"chain_id"`
	RuleId  int    `json:"rule_id"`
}
type RespRuleGetRuleInfo struct {
	util.CommonError
	RuleInfo struct {
		OwnerCorpRange struct {
			DepartmentIds []string `json:"departmentids"`
			UserIds       []string `json:"userids"`
		} `json:"owner_corp_range"`
		MemberCorpRange struct {
			GroupIds []string `json:"groupids"`
			CorpIds  []string `json:"corpids"`
		} `json:"member_corp_range"`
	} `json:"rule_info"`
}

type ReqRuleAddRule struct {
	ChainId  string `json:"chain_id"`
	RuleInfo struct {
		OwnerCorpRange struct {
			DepartmentIds []string `json:"departmentids"`
			UserIds       []string `json:"userids"`
		} `json:"owner_corp_range"`
		MemberCorpRange struct {
			GroupIds []string `json:"groupids"`
			CorpIds  []string `json:"corpids"`
		} `json:"member_corp_range"`
	} `json:"rule_info"`
}
type RespRuleAddRule struct {
	util.CommonError
	RuleId int `json:"rule_id"`
}

type ReqRuleModifyRule struct {
	ChainId  string `json:"chain_id"`
	RuleId   int    `json:"rule_id"`
	RuleInfo struct {
		OwnerCorpRange struct {
			DepartmentIds []string `json:"departmentids"`
			UserIds       []string `json:"userids"`
		} `json:"owner_corp_range"`
		MemberCorpRange struct {
			GroupIds []string `json:"groupids"`
			CorpIds  []string `json:"corpids"`
		} `json:"member_corp_range"`
	} `json:"rule_info"`
}
