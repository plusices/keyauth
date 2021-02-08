package endpoint

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*router.Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
	}
}

// NewDefaultRegistryRequest todo
func NewDefaultRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Entries: []*router.Entry{},
	}
}

// Validate 校验注册请求合法性
func (req *RegistryRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require *router.Entry")
	}

	// tk := req.GetToken()
	// if tk == nil {
	// 	return fmt.Errorf("token required when service endpoints registry")
	// }

	// if !tk.UserType.IsIn(types.UserType_SERVICE) {
	// 	return fmt.Errorf("only service account can registry endpoints")
	// }

	return validate.Struct(req)
}

// Endpoints 功能列表
func (req *RegistryRequest) Endpoints(serviceID string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			Id:        GenHashID(serviceID, req.Entries[i].Path, req.Entries[i].Method),
			CreateAt:  ftime.Now().Timestamp(),
			UpdateAt:  ftime.Now().Timestamp(),
			ServiceId: serviceID,
			Version:   req.Version,
			Entry:     req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

// NewRegistryResponse todo
func NewRegistryResponse(message string) *RegistryResponse {
	return &RegistryResponse{Message: message}
}

// NewQueryEndpointRequestFromHTTP 列表查询请求
func NewQueryEndpointRequestFromHTTP(r *http.Request) *QueryEndpointRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	return &QueryEndpointRequest{
		Page:         &page.PageRequest,
		ServiceId:    qs.Get("service_id"),
		Path:         qs.Get("path"),
		Method:       qs.Get("method"),
		FunctionName: qs.Get("function_name"),
		Resource:     qs.Get("resource"),
	}
}

// NewQueryEndpointRequest 列表查询请求
func NewQueryEndpointRequest(pageReq *request.PageRequest) *QueryEndpointRequest {
	return &QueryEndpointRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewDescribeEndpointRequestWithID todo
func NewDescribeEndpointRequestWithID(id string) *DescribeEndpointRequest {
	return &DescribeEndpointRequest{Id: id}
}

// Validate 校验
func (req *DescribeEndpointRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}

// NewDeleteEndpointRequestWithServiceID todo
func NewDeleteEndpointRequestWithServiceID(id string) *DeleteEndpointRequest {
	return &DeleteEndpointRequest{ServiceId: id}
}

// NewQueryResourceRequestFromHTTP 列表查询请求
func NewQueryResourceRequestFromHTTP(r *http.Request) *QueryResourceRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	return &QueryResourceRequest{
		Page:       &page.PageRequest,
		ServiceIds: strings.Split(qs.Get("service_ids"), ","),
	}
}
