package svc

import (
	"gorm.io/gorm"
	"zero_study/common/init_gorm"
	"zero_study/model_study/user_gorm/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		DB:     init_gorm.InitGorm(c.Mysql.DataSource),
	}
}
