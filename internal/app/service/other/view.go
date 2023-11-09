package other

import "grapefruit/internal/app/service"

var ViewService service.IView = &View{}

type View struct {
}

func (v *View) View() {
	//TODO implement me
	panic("implement me")
	//todo 验证码检查是不是正确，正确才可以继续浏览
}
