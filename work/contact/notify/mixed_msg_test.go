package notify

import (
	"encoding/xml"
	"testing"
)

func TestNotify_UnmarshalMixedMsg(t *testing.T) {
	memberCreateEventRawStr := `<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName> 
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>create_user</ChangeType>
 	<NewUserID><![CDATA[zhangsan001]]></NewUserID>
    <UserID><![CDATA[zhangsan]]></UserID>
    <Name><![CDATA[张三]]></Name>
    <Department><![CDATA[1,2,3]]></Department>
    <MainDepartment>1</MainDepartment>
    <IsLeaderInDept><![CDATA[1,0,0]]></IsLeaderInDept>
    <Position><![CDATA[产品经理]]></Position>
    <Mobile>13800000000</Mobile>
    <Gender>1</Gender>
    <Email><![CDATA[zhangsan@gzdev.com]]></Email>
    <Status>1</Status>
    <Avatar><![CDATA[http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0]]></Avatar>
    <Alias><![CDATA[zhangsan]]></Alias>
    <Telephone><![CDATA[020-123456]]></Telephone>
    <Address><![CDATA[广州市]]></Address>
    <ExtAttr>
        <Item>
        <Name><![CDATA[爱好]]></Name>
        <Type>0</Type>
        <Text>
            <Value><![CDATA[旅游]]></Value>
        </Text>
        </Item>
        <Item>
        <Name><![CDATA[卡号]]></Name>
        <Type>1</Type>
        <Web>
            <Title><![CDATA[企业微信]]></Title>
            <Url><![CDATA[https://work.weixin.qq.com]]></Url>
        </Web>
        </Item>
    </ExtAttr>

    <Id>2</Id>
    <Name><![CDATA[张三]]></Name>
    <ParentId><![CDATA[1]]></ParentId>
    <Order>1</Order>
    <TagId>1</TagId>
    <AddUserItems><![CDATA[zhangsan,lisi]]></AddUserItems>
    <DelUserItems><![CDATA[zhangsan1,lisi1]]></DelUserItems>
    <AddPartyItems><![CDATA[1,2]]></AddPartyItems>
    <DelPartyItems><![CDATA[3,4]]></DelPartyItems>

	<BatchJob>
		<JobId><![CDATA[S0MrnndvRG5fadSlLwiBqiDDbM143UqTmKP3152FZk4]]></JobId>
		<JobType><![CDATA[sync_user]]></JobType>
		<ErrCode>0</ErrCode>
		<ErrMsg><![CDATA[ok]]></ErrMsg>
	</BatchJob>
</xml>`
	var msg MixedMsg
	err := xml.Unmarshal([]byte(memberCreateEventRawStr), &msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(msg)
	t.Log(msg.Member)
	t.Log(msg.Department)
	t.Log(msg.Tag)
	t.Log(msg.MsgHeader)
	t.Log(msg.BatchJob)
}
