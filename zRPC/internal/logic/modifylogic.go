package logic

import (
	"context"

	"github.com/jasonzou/zRPC/internal/svc"
	"github.com/jasonzou/zRPC/pb/zRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyLogic {
	return &ModifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyLogic) Modify(in *zRPC.Entry) (*zRPC.ModifyResponse, error) {
	// todo: add your logic here and delete this line

	return &zRPC.ModifyResponse{}, nil
}
