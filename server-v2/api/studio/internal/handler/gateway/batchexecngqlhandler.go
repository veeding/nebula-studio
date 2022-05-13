// Code generated by goctl. DO NOT EDIT.
package gateway

import (
	"net/http"

	"github.com/vesoft-inc/go-pkg/validator"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/pkg/ecode"

	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/logic/gateway"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/svc"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BatchExecNGQLHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchExecNGQLParams
		if err := httpx.Parse(r, &req); err != nil {
			err = ecode.WithCode(ecode.ErrParam, err)
			svcCtx.ResponseHandler.Handle(w, r, nil, err)
			return
		}
		if err := validator.Struct(req); err != nil {
			svcCtx.ResponseHandler.Handle(w, r, nil, err)
			return
		}

		l := gateway.NewBatchExecNGQLLogic(r.Context(), svcCtx)
		data, err := l.BatchExecNGQL(req)
		svcCtx.ResponseHandler.Handle(w, r, data, err)
	}
}
