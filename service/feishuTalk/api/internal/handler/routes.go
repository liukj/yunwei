// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	report "ywadmin-v3/service/feishuTalk/api/internal/handler/report"
	"ywadmin-v3/service/feishuTalk/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/feishuEvent",
				Handler: report.InfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}