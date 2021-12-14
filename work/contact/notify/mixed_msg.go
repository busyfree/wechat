package notify

type (
	MsgType    string
	EventType  string
	ChangeType string
)

// MsgHeader 推送过来的消息(事件)的通用消息头
type MsgHeader struct {
	ToUserName   string  `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"   json:"CreateTime"`
	MsgType      MsgType `xml:"MsgType"      json:"MsgType"`
}

// MixedMsg 通讯录回调推送过来的消息(事件)的合集
// https://work.weixin.qq.com/api/doc/90000/90135/90967
type MixedMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader
	EventType  EventType  `xml:"Event" json:"Event"`
	ChangeType ChangeType `xml:"ChangeType" json:"ChangeType"`
	Name       string     `xml:"Name,omitempty" json:"Name,omitempty"`
	BatchJob   *BatchJob  `xml:"BatchJob,omitempty" json:"BatchJob,omitempty"`
	*Member
	*Department
	*Tag
}

type Member struct {
	NewUserID      string `xml:"NewUserID,omitempty" json:"NewUserID,omitempty"`
	UserID         string `xml:"UserID" json:"UserID"`
	Department     string `xml:"Department" json:"Department"`
	MainDepartment string `xml:"MainDepartment" json:"MainDepartment"`
	IsLeaderInDept string `xml:"IsLeaderInDept" json:"IsLeaderInDept"`
	Position       string `xml:"Position" json:"Position"`
	Mobile         string `xml:"Mobile" json:"Mobile"`
	Gender         string `xml:"Gender" json:"Gender"`
	Email          string `xml:"Email" json:"Email"`
	Status         string `xml:"Status" json:"Status"`
	Avatar         string `xml:"Avatar" json:"Avatar"`
	Alias          string `xml:"Alias" json:"Alias"`
	Telephone      string `xml:"Telephone" json:"Telephone"`
	Address        string `xml:"Address" json:"Address"`
	ExtAttr        struct {
		Item []struct {
			Name string `xml:"Name" json:"Name"`
			Type string `xml:"Type" json:"Type"`
			Text struct {
				Value string `xml:"Value" json:"Value"`
			} `xml:"Text" json:"Text,omitempty"`
			Web struct {
				Title string `xml:"Title" json:"Title"`
				Url   string `xml:"Url" json:"Url"`
			} `xml:"Web" json:"Web,omitempty"`
		} `xml:"Item" json:"Item"`
	} `xml:"ExtAttr" json:"ExtAttr"`
}

type Department struct {
	Id       string `xml:"Id" json:"Id"`
	ParentId string `xml:"ParentId" json:"ParentId"`
	Order    string `xml:"Order" json:"Order"`
}

type Tag struct {
	TagId         string `xml:"TagId" json:"TagId"`
	AddUserItems  string `xml:"AddUserItems" json:"AddUserItems"`
	DelUserItems  string `xml:"DelUserItems" json:"DelUserItems"`
	AddPartyItems string `xml:"AddPartyItems" json:"AddPartyItems"`
	DelPartyItems string `xml:"DelPartyItems" json:"DelPartyItems"`
}

type BatchJob struct {
	JobId   string `xml:"JobId" json:"JobId"`
	JobType string `xml:"JobType" json:"JobType"`
	ErrCode string `xml:"ErrCode" json:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg" json:"ErrMsg"`
}
