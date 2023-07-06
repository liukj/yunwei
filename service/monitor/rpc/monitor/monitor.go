// Code generated by goctl. DO NOT EDIT.
// Source: monitor.proto

package monitor

import (
	"context"

	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GraphListReq     = monitorclient.GraphListReq
	GraphListResp    = monitorclient.GraphListResp
	ReportAddReq     = monitorclient.ReportAddReq
	ReportAddResp    = monitorclient.ReportAddResp
	SelectReportReq  = monitorclient.SelectReportReq
	SelectReportResp = monitorclient.SelectReportResp

	Monitor interface {
		ReportAdd(ctx context.Context, in *ReportAddReq, opts ...grpc.CallOption) (*ReportAddResp, error)
		GraphList(ctx context.Context, in *GraphListReq, opts ...grpc.CallOption) (*GraphListResp, error)
		SelectReport(ctx context.Context, in *SelectReportReq, opts ...grpc.CallOption) (*SelectReportResp, error)
	}

	defaultMonitor struct {
		cli zrpc.Client
	}
)

func NewMonitor(cli zrpc.Client) Monitor {
	return &defaultMonitor{
		cli: cli,
	}
}

func (m *defaultMonitor) ReportAdd(ctx context.Context, in *ReportAddReq, opts ...grpc.CallOption) (*ReportAddResp, error) {
	client := monitorclient.NewMonitorClient(m.cli.Conn())
	return client.ReportAdd(ctx, in, opts...)
}

func (m *defaultMonitor) GraphList(ctx context.Context, in *GraphListReq, opts ...grpc.CallOption) (*GraphListResp, error) {
	client := monitorclient.NewMonitorClient(m.cli.Conn())
	return client.GraphList(ctx, in, opts...)
}

func (m *defaultMonitor) SelectReport(ctx context.Context, in *SelectReportReq, opts ...grpc.CallOption) (*SelectReportResp, error) {
	client := monitorclient.NewMonitorClient(m.cli.Conn())
	return client.SelectReport(ctx, in, opts...)
}