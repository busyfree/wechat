package notify

import (
	"github.com/silenceper/wechat/v2/work/context"
)

// Notify 微信客服实例
type Notify struct {
	ctx *context.Context
}

// NewNotify 初始化通讯录回调实例
func NewNotify(ctx *context.Context) (client *Notify) {
	client = &Notify{
		ctx: ctx,
	}
	return client
}
