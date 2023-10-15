package service

type IChatService interface {
	RunAsClient()
	RunAsServer()
	Start()
	Stop()
}
