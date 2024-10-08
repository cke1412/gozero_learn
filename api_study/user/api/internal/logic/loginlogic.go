package logic

import (
	"context"

	"zero_study/api_study/user/api/internal/svc"
	"zero_study/api_study/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	return &types.LoginResponse{
		Code: 0,
		Data: "xxx.xxx.xxx",
		Msg:  "Success",
	}, nil
}
