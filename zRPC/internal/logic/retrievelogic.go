package logic

import (
	"context"

	"github.com/jasonzou/zRPC/internal/svc"
	"github.com/jasonzou/zRPC/pb/zRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveLogic {
	return &RetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveLogic) Retrieve(in *zRPC.RetrieveRequest) (*zRPC.Entry, error) {
	// todo: add your logic here and delete this line

	return &zRPC.Entry{}, nil
}
