package logic

import (
	"context"

	"zero_study/api_study/user/api/internal/svc"
	"zero_study/api_study/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {

	return &types.UserInfoResponse{
		Code: 0,
		Data: types.UserInfo{
			UserId:   1,
			Username: "ckck",
		},
		Msg: "Succcess",
	}, nil
}
