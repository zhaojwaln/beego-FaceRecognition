package service

import (
	"io/ioutil"
	"beego-FaceRecognition/src/common"
	"net/http"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)


var BAIDU_AI_APPID = beego.AppConfig.String("app_id")
var BAIDU_AI_KEY = beego.AppConfig.String("client_id")
var BAIDU_AI_CRET = beego.AppConfig.String("client_secret")
var TOKEN_API_URL = beego.AppConfig.String("token_api_url")
var FACEDETECT_API_URL = beego.AppConfig.String("facedetect_url")
var FACEIDENTIFY_API_URL = beego.AppConfig.String("faceidentify_url")
var FACEMATCH_API_URL = beego.AppConfig.String("facematch_url")

type faceService struct{
	Client *http.Client
	accessToken string
}

func (this *faceService) GetInterFaceUrl(baseurl string) string {
	this.Client = &http.Client{}
	if this.accessToken == "" {
		doflag, _ := this.getToken()

		if !doflag {
			panic("获取accessToken失败")
		}
	}

	return fmt.Sprintf("%s?access_token=%s", baseurl, this.accessToken)
}

func (this *faceService) getToken() (bool, string) {
	postArg := map[string] interface{}{
		"client_id": 	 BAIDU_AI_KEY,
		"client_secret": BAIDU_AI_CRET,
		"grant_type": 	 "client_credentials",
	}
	
	resp, _ := this.Client.PostForm(TOKEN_API_URL, common.InitPostData(postArg))
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	mapResult := make(map[string]interface{})
	json.Unmarshal(data, &mapResult)

	accessToken, ok := mapResult["access_token"]

	if ok {
		this.accessToken = accessToken.(string)
		return true, this.accessToken
	}

	return false, ""
}