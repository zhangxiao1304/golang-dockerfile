package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppInfoController struct {
}

func (AppInfoController) AppInfo(c *gin.Context) {
	resp := make(map[string]interface{})
	resp["code"] = 200
	resp["msg"] = "hello golang-docker"

	c.JSON(http.StatusOK, resp)
}
