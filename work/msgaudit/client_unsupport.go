//go:build !linux || !cgo || !msgaudit
// +build !linux !cgo !msgaudit

// Package msgaudit for unsupport platform
package msgaudit

import (
	"fmt"

	"github.com/silenceper/wechat/v2/work/context"
)

// Client 会话存档
type Client struct {
}

// NewClient new
func NewClient(ctx *context.Context) (*Client, error) {
	return nil, fmt.Errorf("会话存档功能目前只支持Linux平台运行，并且打开设置CGO_ENABLED=1")
}
