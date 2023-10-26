package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grapefruit/config"
	"grapefruit/internal/app/service/s3buketservice"
	"grapefruit/internal/domain/do"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Service struct {
	MiniProgramAccessToken string
	MiniProgramExpiredTime time.Time
	mutex                  sync.Mutex
}

func (w *Service) GetAccessToken() (string, error) {
	duration := time.Since(w.MiniProgramExpiredTime)
	if duration > 0 {
		w.fetchAccessToken(config.MiniProgramFetchAccessTokenTimes)
	}
	return w.MiniProgramAccessToken, nil
}

func (w *Service) fetchAccessToken(tryTimes int) string {
	for i := 0; i < tryTimes; i++ {
		resp, err := http.Get(config.FetchMiniProgramAccessTokenCompleteUrl)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()
		ret := do.WechatAccessTokenResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return ""
		}

		w.MiniProgramAccessToken = ret.AccessToken
		w.MiniProgramExpiredTime = time.Now().Add(time.Second * time.Duration(ret.ExpiresIn))
		return w.MiniProgramAccessToken
	}
	return ""
}

func (w *Service) GetMiniProgramQRCode(req *do.ClientCreateQRCodeReq) (url string, err error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return "", err
	}
	reqUrl := fmt.Sprintf(config.Getwxacodeunlimit, accessToken)
	param := &do.QueryQRCodeReq{
		Path:  fmt.Sprintf("pages/index/index"),
		Width: 430,
		Scene: "",
	}

	bytesStream, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", reqUrl, bytes.NewReader(bytesStream))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {

		}
	}()
	saveQRCodeKey := config.MakeS3SaveQRCodeKey(0, "")
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(saveQRCodeKey, dataBytes, 0666)
	if err != nil {
		return "", err
	}
	err = s3buketservice.BucketBasicsService.UploadFile(" ", saveQRCodeKey, saveQRCodeKey, "image/png")
	if err != nil {
		return
	}
	url = config.S3BucketUrl + saveQRCodeKey
	err = os.Remove(saveQRCodeKey)
	return url, err
}
