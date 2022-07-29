package oa

import (
	"github.com/silenceper/wechat/v2/work/context"
	"github.com/silenceper/wechat/v2/work/oa/checkin"
)

type OA struct {
	ctx *context.Context
}

func NewOA(ctx *context.Context) *OA {
	return &OA{
		ctx,
	}
}

func (c *OA) GetCheckIn() *checkin.OACheckIn {
	return checkin.NewOACheckIn(c.ctx)
}
