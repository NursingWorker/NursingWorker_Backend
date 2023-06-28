package weixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	OpenID      string `json:"openid"`
}

type UserInfoResponse struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"headimgurl"`
}

// 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, viper.GetString("AppID"), viper.GetString("AppSecret"), code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	return &wxResp, nil
}

func GetUserInfo(accessToken string, openID string) (string, string, error) {
	userInfoURL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", accessToken, openID)
	response, err := http.Get(userInfoURL)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	var userInfoResponse UserInfoResponse
	err = json.Unmarshal(body, &userInfoResponse)
	if err != nil {
		return "", "", err
	}

	return userInfoResponse.Nickname, userInfoResponse.Avatar, nil
}

func GetAccessToken(code string) (string, string, error) {
	accessTokenURL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", viper.GetString("AppID"), viper.GetString("AppSecret"), code)
	response, err := http.Get(accessTokenURL)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return "", "", err
	}

	return accessTokenResponse.AccessToken, accessTokenResponse.OpenID, nil
}
