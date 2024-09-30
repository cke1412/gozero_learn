package logic

import (
	"context"
	"errors"
	"zero_study/model_study/user/api/internal/svc"
	"zero_study/model_study/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	//res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
	//	Username: "ckckck",
	//	Password: "123456",
	//})
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println(res)

	user, err := l.svcCtx.UserModel.FindOneByUsernameAndPassword(l.ctx, req.Username, req.Password)
	if err != nil {
		return "", errors.New("login failed")
	}

	return user.Username, nil
}
