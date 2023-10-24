package login

import (
	"context"
	"gorm.io/gorm"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/app/service"
	"grapefruit/internal/domain/po"
)

func GetUserTable() (tx *gorm.DB) {
	tableUser := psql.GetGormDB().Table("user")
	return tableUser
}

var LoginService service.ILogin = &_Login{}

type _Login struct {
}

func (l *_Login) Login(name, pwd string) error {
	tableUser := GetUserTable()
	user := &po.User{}
	tx := tableUser.WithContext(context.Background()).Where("name=? and pwd=?", name, pwd).Find(user)
	return tx.Error
}

func (l *_Login) Register(name, pwd string) error {
	tableUser := GetUserTable()
	user := &po.User{Name: name, Pwd: pwd}
	tx := tableUser.WithContext(context.Background()).Clauses().Create(user)
	return tx.Error
}