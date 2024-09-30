package init_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zero_study/model_study/user_gorm/models"
)

func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("Mysql connection failed, error=" + err.Error())
	}
	fmt.Println("Mysql connection has been established")
	db.AutoMigrate(&models.UserModel{})

	return db
}
