package checkin

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/xerror"
)

const (
	getHardwareCheckinData = "/cgi-bin/hardware/get_hardware_checkin_data?access_token=%s"
	getCheckinData         = "/cgi-bin/hardware/getcheckindata?access_token=%s"
)

type OACheckIn struct {
	*context.Context
}

func NewOACheckIn(ctx *context.Context) *OACheckIn {
	return &OACheckIn{
		ctx,
	}
}

// GetHardwareCheckinData 获取设备打卡数据
// https://developer.work.weixin.qq.com/document/path/94126
func (r *OACheckIn) GetHardwareCheckinData(options GetHardwareCheckinDataReq) (info GetHardwareCheckinDataResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(r.GetQYAPIDomain()+fmt.Sprintf(getHardwareCheckinData, accessToken), options)
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

// GetCheckinData 获取打卡记录数据
// https://developer.work.weixin.qq.com/document/path/90262
func (r *OACheckIn) GetCheckinData(options GetCheckinDataReq) (info GetCheckinDataResp, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(r.GetQYAPIDomain()+fmt.Sprintf(getCheckinData, accessToken), options)
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
