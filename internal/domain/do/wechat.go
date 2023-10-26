package do

//QueryQRCodeReq
//{
// "page": "pages/index/index",
// "scene": "a=1",
// "check_path": true,
// "env_version": "release"
//}

type QueryQRCodeReq struct {
	Path  string `json:"path"`
	Width int    `json:"width"`
	Scene string `json:"scene"`
}

type ClientCreateQRCodeReq struct {
}

type WechatAccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}
