package oa

import (
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/oa/checkin"
)

type OA struct {
	*context.Context
}

func NewOA(ctx *context.Context) *OA {
	return &OA{
		ctx,
	}
}

func (c *OA) GetCheckIn(ctx *context.Context) *checkin.OACheckIn {
	return checkin.NewOACheckIn(ctx)
}
