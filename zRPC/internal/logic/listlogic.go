package logic

import (
	"context"

	"github.com/jasonzou/zRPC/internal/svc"
	"github.com/jasonzou/zRPC/pb/zRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *zRPC.ListRequest) (*zRPC.Entries, error) {
	// todo: add your logic here and delete this line

	return &zRPC.Entries{}, nil
}
