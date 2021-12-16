package kf

import (
	"github.com/silenceper/wechat/v2/work/context"
)

// Client 微信客服实例
type Client struct {
	ctx *context.Context
}

// NewClient 初始化微信客服实例
func NewClient(ctx *context.Context) (client *Client) {
	client = &Client{
		ctx: ctx,
	}
	return client
}
