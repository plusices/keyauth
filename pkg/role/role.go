package role

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
)

const (
	// MaxPermissionCount 一个角色最多可以容纳的权限条数
	MaxPermissionCount = 500
)

// New 新创建一个Role
func New(t Type, req *CreateRoleRequest) (*Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Role{
		Type:              t,
		CreateAt:          ftime.Now(),
		UpdateAt:          ftime.Now(),
		CreateRoleRequest: req,
	}, nil
}

// NewDefaultRole 默认实例
func NewDefaultRole() *Role {
	return &Role{
		CreateRoleRequest: NewCreateRoleRequest(),
	}
}

// Role is rbac's role
type Role struct {
	ID                 string     `bson:"_id" json:"id"`                        // 角色ID
	Type               Type       `bson:"type" json:"type"`                     // 角色类型
	CreateAt           ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间`
	UpdateAt           ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	*CreateRoleRequest `bson:",inline"`
}

// NewCreateRoleRequest 实例化请求
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{}
}

// CreateRoleRequest 创建应用请求
type CreateRoleRequest struct {
	Name        string        `bson:"name" json:"name,omitempty" validate:"required,lte=30"`       // 应用名称
	Description string        `bson:"description" json:"description,omitempty" validate:"lte=400"` // 应用简单的描述
	Permissions []*Permission `bson:"permissions" json:"permissions,omitempty"`                    // 读权限
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	pc := len(req.Permissions)
	if pc > MaxPermissionCount {
		return fmt.Errorf("role permission overed max count: %d",
			MaxPermissionCount)
	}
	return validate.Struct(req)
}

// CheckPermission 检测该角色是否具有该权限
func (r *Role) CheckPermission() error {
	return nil
}

// NewRoleSet 实例化
func NewRoleSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*Role{},
	}
}

// Set 角色集合
type Set struct {
	*request.PageRequest

	Total int64   `json:"total"`
	Items []*Role `json:"items"`
}

// Add todo
func (s *Set) Add(item *Role) {
	s.Items = append(s.Items, item)
}

// Permission 权限
type Permission struct {
	Effect       EffectType `bson:"effect" json:"effect,omitempty"`               // 效力
	ResourceName string     `bson:"resource_name" json:"resource_name,omitempty"` // 资源列表
	LabelKey     string     `bson:"label_key" json:"label_key,omitempty"`
	LabelValues  []string   `bson:"label_values" json:"label_values,omitempty"` // 标识
}

// ID 计算唯一ID
func (p *Permission) ID(namespace string) string {
	return namespace + "." + p.ResourceName
}

// PermissionSet 用户列表
type PermissionSet struct {
	*request.PageRequest

	Total int64         `json:"total"`
	Items []*Permission `json:"items"`
}

// Add todo
func (s *PermissionSet) Add(item *Permission) {
	s.Items = append(s.Items, item)
}
