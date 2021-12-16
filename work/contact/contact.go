package contact

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/contact/department"
	"github.com/silenceper/wechat/v2/work/contact/linkedcorp"
	"github.com/silenceper/wechat/v2/work/contact/member"
	"github.com/silenceper/wechat/v2/work/contact/notify"
	"github.com/silenceper/wechat/v2/work/contact/tag"
	"github.com/silenceper/wechat/v2/work/context"
)

// Contact 企业微信通讯录相关API
type Contact struct {
	ctx *context.Context
}

// NewContact 实例化企业微信通讯录API
func NewContact(ctx *context.Context) (client *Contact) {
	client = &Contact{
		ctx: ctx,
	}
	return client
}

// SetAccessTokenHandle 自定义access_token获取方式
func (r *Contact) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	r.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (r *Contact) GetContext() *context.Context {
	return r.ctx
}

// GetDepartment 返回 department 实例
func (r *Contact) GetDepartment() *department.Department {
	return department.NewDepartment(r.ctx)
}

// GetLinkedCorp 返回 linked corp 实例
func (r *Contact) GetLinkedCorp() *linkedcorp.LinkedCorp {
	return linkedcorp.NewLinkedCorp(r.ctx)
}

// GetMember 返回 member 实例
func (r *Contact) GetMember() *member.Member {
	return member.NewMember(r.ctx)
}

// GetTag 返回 tag 实例
func (r *Contact) GetTag() *tag.Tag {
	return tag.NewTag(r.ctx)
}

// GetNotify 返回 notify 实例
func (r *Contact) GetNotify() *notify.Notify {
	return notify.NewNotify(r.ctx)
}
