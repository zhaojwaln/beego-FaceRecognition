package main

import (
	"github.com/astaxie/beego"

	_ "beego-FaceRecognition/src/routers"
	_ "beego-FaceRecognition/src/service"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("console", "")
	beego.Run()
}
