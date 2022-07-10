package router

import (
	"chatgo/imsrv/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	ApiRouter := gin.Default()
	ApiRouter.LoadHTMLGlob("public/html/*")
	ApiRouter.StaticFS("/public", http.Dir("./public"))

	ApiRouter.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", struct{}{})
	})

	ApiRouter.GET("/chat", func(c *gin.Context) {
		data := gin.H{
			"title": "chatgo",
		}
		c.HTML(http.StatusOK, "chat.html", data)
	})

	ApiRouter.GET("/ws", api.Ws)

	// ApiGroup := ApiRouter.Group("/v1")

	// router.InitUserRouter(ApiGroup)

	return ApiRouter
}
