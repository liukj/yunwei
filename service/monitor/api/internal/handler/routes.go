// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	graph "ywadmin-v3/service/monitor/api/internal/handler/graph"
	report "ywadmin-v3/service/monitor/api/internal/handler/report"
	"ywadmin-v3/service/monitor/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: report.ReportAddHandler(serverCtx),
			},
		},
		rest.WithPrefix("/monitor/report"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: graph.GraphListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/monitor/graph"),
	)
}
