package main

import (
	"github.com/gin-gonic/gin"
)

var startMode = gin.DebugMode

func main() {
	println("[Server Starting]...")

	gin.SetMode(startMode)

	//注册路由
	//router := routers.InitRouter()
}
