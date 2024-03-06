package router

import (
	"github.com/gin-gonic/gin"
	"golang-run-dockerfile/internal/controller"
)

func InitRouter(engine *gin.Engine) {
	ApiDocs(engine)
}

func ApiDocs(engine *gin.Engine) {

	apiDocs := engine

	apiGourp := apiDocs.Group("/inner")
	infoController := controller.AppInfoController{}

	{
		apiGourp.GET("/appInfo", infoController.AppInfo)
	}

}
