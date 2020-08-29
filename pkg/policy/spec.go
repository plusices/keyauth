package policy

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 策略服务
type Service interface {
	CreatePolicy(req *CreatePolicyRequest) (*Policy, error)
	QueryPolicy(req *QueryPolicyRequest) (*Set, error)
	DescribePolicy(req *DescribePolicyRequest) (*Policy, error)
}

// NewQueryPolicyRequestFromHTTP 列表查询请求
func NewQueryPolicyRequestFromHTTP(r *http.Request) *QueryPolicyRequest {
	page := request.NewPageRequestFromHTTP(r)
	req := &QueryPolicyRequest{
		Session:     token.NewSession(),
		PageRequest: page,
	}

	qs := r.URL.Query()
	req.Account = qs.Get("account")
	req.RoleID = qs.Get("role_id")
	req.NamespaceID = qs.Get("namespace_id")
	return req
}

// NewQueryPolicyRequest 列表查询请求
func NewQueryPolicyRequest(pageReq *request.PageRequest) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Session:     token.NewSession(),
		PageRequest: pageReq,
	}
}

// QueryPolicyRequest 获取子账号列表
type QueryPolicyRequest struct {
	*request.PageRequest
	*token.Session

	Account     string `json:"account,omitempty"`
	RoleID      string `json:"role_id,omitempty"`
	NamespaceID string `json:"namespace_id,omitempty"`
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{}
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

// NewDescriptPolicyRequest new实例
func NewDescriptPolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{
		Session: token.NewSession(),
	}
}

// DescribePolicyRequest todo
type DescribePolicyRequest struct {
	*token.Session
	ID string
}

// Validate todo
func (req *DescribePolicyRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}
