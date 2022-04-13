// Code generated by goctl. DO NOT EDIT!
// Source: header.proto

package server

import (
	"context"

	"github.com/jasonzou/zRPC/internal/logic"
	"github.com/jasonzou/zRPC/internal/svc"
	"github.com/jasonzou/zRPC/pb/zRPC"
)

type HeaderServiceServer struct {
	svcCtx *svc.ServiceContext
	zRPC.UnimplementedHeaderServiceServer
}

func NewHeaderServiceServer(svcCtx *svc.ServiceContext) *HeaderServiceServer {
	return &HeaderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HeaderServiceServer) Add(ctx context.Context, in *zRPC.Header) (*zRPC.AddHeaderResponse, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *HeaderServiceServer) Modify(ctx context.Context, in *zRPC.Header) (*zRPC.ModifyHeaderResponse, error) {
	l := logic.NewModifyLogic(ctx, s.svcCtx)
	return l.Modify(in)
}

func (s *HeaderServiceServer) Retrieve(ctx context.Context, in *zRPC.RetrieveHeaderRequest) (*zRPC.Header, error) {
	l := logic.NewRetrieveLogic(ctx, s.svcCtx)
	return l.Retrieve(in)
}