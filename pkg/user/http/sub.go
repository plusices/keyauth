package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (h *handler) CreateSubAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	// 非管理员, 主账号 可以创建子账号
	if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_INTERNAL, types.UserType_PRIMARY) {
		response.Failed(w, exception.NewPermissionDeny("%s user can't create sub account", tk.UserType))
		return
	}

	req := user.NewCreateUserRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	req.UserType = types.UserType_SUB
	req.CreateType = user.CreateType_DOMAIN_CREATED

	d, err := h.service.CreateAccount(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) QuerySubAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := user.NewNewQueryAccountRequestFromHTTP(r)
	req.UserType = types.UserType_SUB
	req.Domain = tk.Domain

	d, err := h.service.QueryAccount(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) DescribeSubAccount(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := user.NewDescriptAccountRequestWithAccount(rctx.PS.ByName("account"))

	d, err := h.service.DescribeAccount(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) PatchSubAccount(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := user.NewPatchAccountRequest()
	req.Account = rctx.PS.ByName("account")

	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateAccountProfile(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

// DestroySubAccount 注销账号
func (h *handler) DestroySubAccount(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := &user.DeleteAccountRequest{Account: rctx.PS.ByName("account")}

	_, err := h.service.DeleteAccount(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}

func (h *handler) BlockSubAccount(w http.ResponseWriter, r *http.Request) {
	req := user.NewBlockAccountRequest("", "")
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.BlockAccount(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
}

func (h *handler) UnBlockSubAccount(w http.ResponseWriter, r *http.Request) {
	req := user.NewBlockAccountRequest("", "")
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.BlockAccount(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
}
