/*
 * Copyright  (c) 2022 MS. All rights reserved.
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
 *
 * File:    def.go
 * Created: 2022/7/29 10:47:29
 * Authors: MS<geek.snail@qq.com>
 */

package checkin

import (
	"github.com/silenceper/wechat/v2/util"
)

type GetHardwareCheckinDataReq struct {
	FilterType *int     `json:"filter_type"`
	Starttime  int64    `json:"starttime"`
	Endtime    int64    `json:"endtime"`
	UserIdList []string `json:"useridlist"`
}

type GetHardwareCheckinDataResp struct {
	util.CommonError
	List []struct {
		UserId         string `json:"userid"`
		CheckinTime    int64  `json:"checkin_time"`
		DeviceSerialNo string `json:"device_sn"`
		DeviceName     string `json:"device_name"`
	} `json:"checkindata"`
}

type GetCheckinDataReq struct {
	CheckinType int      `json:"opencheckindatatype"`
	StartTime   int64    `json:"starttime"`
	EndTime     int64    `json:"endtime"`
	UserIdList  []string `json:"useridlist"`
}

type GetCheckinDataResp struct {
	util.CommonError
	List []struct {
		UserId              string   `json:"userid"`
		GroupName           string   `json:"groupname"`
		CheckinType         string   `json:"checkin_type"`
		ExceptionType       string   `json:"exception_type"`
		CheckinTime         int64    `json:"checkin_time"`
		LocationTitle       string   `json:"location_title"`
		LocationDetail      string   `json:"location_detail"`
		WifiName            string   `json:"wifiname"`
		Notes               string   `json:"notes"`
		WifiMac             string   `json:"wifimac"`
		MediaIds            []string `json:"mediaids"`
		ScheduleCheckinTime int64    `json:"sch_checkin_time"`
		GroupId             int      `json:"groupid"`
		ScheduleId          int      `json:"schedule_id"`
		TimelineId          int      `json:"timeline_id"`
		Latitude            int      `json:"lat,omitempty"`
		Logitude            int      `json:"lng,omitempty"`
		Deviceid            string   `json:"deviceid,omitempty"`
	} `json:"checkindata"`
}
