/*
 * Copyright  (c) 2022 MS. All rights reserved.
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
 *
 * File:    oa.go
 * Created: 2022/7/29 10:27:37
 * Authors: MS<geek.snail@qq.com>
 */

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
