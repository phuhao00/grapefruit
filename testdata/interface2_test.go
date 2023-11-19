package testdata

import (
	"fmt"
	"testing"
)

// 定义一个基础接口
type Messenger interface {
	SendMessage() string
}

// 定义另一个接口，嵌套了基础接口
type AdvancedMessenger interface {
	Messenger // 嵌套 Messenger 接口
	ReceiveMessage() string
}

// 实现基础接口的类型
type Email struct {
	Subject string
	Body    string
}

func (e Email) SendMessage() string {
	return fmt.Sprintf("Sending email: Subject - %s, Body - %s", e.Subject, e.Body)
}

// 实现包含嵌套接口的类型
type Slack struct {
	Channel string
	Message string
}

func (s Slack) SendMessage() string {
	return fmt.Sprintf("Sending Slack message to %s: %s", s.Channel, s.Message)
}

func (s Slack) ReceiveMessage() string {
	return fmt.Sprintf("Receiving Slack message from %s: %s", s.Channel, s.Message)
}

func TestQianTao(t *testing.T) {
	// 使用基础接口
	var email Messenger = Email{Subject: "Hello", Body: "How are you?"}
	fmt.Println(email.SendMessage())

	// 使用嵌套接口
	var slack AdvancedMessenger = Slack{Channel: "#general", Message: "Hello, everyone!"}
	fmt.Println(slack.SendMessage())
	fmt.Println(slack.ReceiveMessage())
}
