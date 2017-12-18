package controllers

import (
	"encoding/json"
	"beego-FaceRecognition/src/common"
	"beego-FaceRecognition/src/service"
	"fmt"
	"encoding/base64"
	"io/ioutil"
	"strings"
	"bytes"
)

type FaceController struct {
	BaseController
}
/**
人脸比对
*/
func (this *FaceController) MatchShow() {
	this.show("face/faceMatch.html")
}

/**
人脸识别
*/
func (this *FaceController) IdentifyShow() {
	this.show("face/faceIdentify.html")
}

/**
人脸检测
*/
func (this *FaceController) DetectShow() {
	this.show("face/faceDetect.html")
}

/**
人脸检测
*/
func (this *FaceController) FaceDetect() {
	pic := this.GetString("pic")

	/*通过百度云api请求人脸检测*/

	idx := strings.IndexAny(pic, ",")
	pic = pic[idx + 1: len(pic)]

	postArg := map[string]interface{}{
		"image": pic,
		"max_face_num": "1",	// 最多处理人脸数目，默认值1
		"face_fields":  "age,beauty,expression,faceshape,gender,glasses,race,qualities",
	}
	
	resp, _ := service.FaceService.Client.PostForm(service.FaceService.GetInterFaceUrl(service.
									FACEDETECT_API_URL), common.InitPostData(postArg))
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		this.jsonResult(err.Error())
	}

	mapResult := make(map[string] interface{})
	json.Unmarshal(data, &mapResult)

	this.jsonResult(mapResult)
}

/**
人脸比对
*/
func (this *FaceController) FaceMatch() {
	picone := this.GetString("picone")
	pictwo := this.GetString("pictwo")
	
	idx := strings.IndexAny(picone, ",")
	picone = picone[idx + 1: len(picone)]
	
	idx = strings.IndexAny(pictwo, ",")
	pictwo = pictwo[idx + 1: len(pictwo)]

	var imagesBlist []string
	imagesBlist = append(imagesBlist, picone)
	imagesBlist = append(imagesBlist, pictwo)

	postArg := map[string]interface{}{
		"images": strings.Join(imagesBlist, ","),
	}

	resp, _ := service.FaceService.Client.PostForm(service.FaceService.GetInterFaceUrl(service.
		FACEMATCH_API_URL), common.InitPostData(postArg))
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		this.jsonResult(err.Error())
	}
	// fmt.Println(string(body))

	mapResult := make(map[string]interface{})
	json.Unmarshal(data, &mapResult)

	fmt.Println(mapResult)

	this.jsonResult(mapResult)
}

/**
人脸识别
*/
func (this *FaceController) FaceIdentify() {
	pic := this.GetString("pic")
	idx := strings.IndexAny(pic, ",")
	pic = pic[idx + 1: len(pic)]
	
	postArg := map[string]interface{}{
		"image": pic,
		"group_id": "group1",
		"user_top_num": 1,
	}

	resp, _ := service.FaceService.Client.PostForm(service.FaceService.GetInterFaceUrl(service.
		FACEIDENTIFY_API_URL), common.InitPostData(postArg))
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		this.jsonResult(err.Error())
	}
	
	mapResult := make(map[string]interface{})

	json.Unmarshal(data, &mapResult)
	fmt.Println(mapResult)

	result := (mapResult["result"].([]interface {}))[0].(map[string]interface{})
	
	var identifyName bytes.Buffer
	identifyName.WriteString("picture/")
	identifyName.WriteString(result["uid"].(string))
	identifyName.WriteString(".jpg")
	fmt.Println(identifyName.String())
	readone, err := ioutil.ReadFile(identifyName.String())
	if nil != err {
		this.jsonResult(err.Error())
	}

	var picBuffer bytes.Buffer
	picBuffer.WriteString("data:image/jpeg;base64,")
	picBuffer.WriteString(base64.StdEncoding.EncodeToString(readone))
	mapResult["result_pic"] = picBuffer.String()

	this.jsonResult(mapResult)
}