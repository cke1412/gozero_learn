package logic

import (
	"context"
	"fmt"
	"zero_study/model_study/user_gorm/api/internal/svc"
	"zero_study/model_study/user_gorm/api/internal/types"
	"zero_study/model_study/user_gorm/models"

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
	err = l.svcCtx.DB.Create(&models.UserModel{
		Username: "ckckck",
		Password: "123456",
	}).Error
	fmt.Println(err)
	return "", nil

	//var user models.UserModel
	//err = l.svcCtx.DB.Take(&user, "user_name = ? and password = ?", req.Username, req.Password).Error
	//if err != nil {
	//	return "", errors.New("login failed")
	//}
	//return user.UserName, nil
}
