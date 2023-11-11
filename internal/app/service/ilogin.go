package service

type ILogin interface {
	Login(name, pwd string) error
	Register(name, pwd, email string) error
}
