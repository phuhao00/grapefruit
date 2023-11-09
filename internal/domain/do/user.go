package do

type UserBaseInfo struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Pwd         string `json:"pwd"`
	Token       string `json:"token"`
	Category    string `json:"category"`
	Email       string `json:"email"`
	Status      int32  `json:"status"`
	Role        int32  `json:"role"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateUserBaseInfoReq struct {
	UserBaseInfo
}

type UpdateUserPhotosReq struct {
	UserId int32    `json:"user_id"`
	Photos []string `json:"photos"`
}

type UpdateUserVlogsReq struct {
	UserId int32    `json:"user_id"`
	Vlogs  []string `json:"vlogs"`
}

type ShareUserIntroductionReq struct {
}
