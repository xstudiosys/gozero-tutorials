package logic

import (
	"context"

	"github.com/jasonzou/zRPC/internal/svc"
	"github.com/jasonzou/zRPC/pb/zRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *zRPC.Header) (*zRPC.AddHeaderResponse, error) {
	// todo: add your logic here and delete this line

	return &zRPC.AddHeaderResponse{}, nil
}
