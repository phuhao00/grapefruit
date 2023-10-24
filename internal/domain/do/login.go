package do

type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type LoginRsp struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type RegisterReq struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type RegisterRsp struct {
}
