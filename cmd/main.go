package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-run-dockerfile/internal/router"
	"log"
	"net/http"
	"os"
)

func main() {
	env := flag.String("env", "", "请输入运行环境:\n uat:本地环境\n test:测试环境\n  prod:正式环境\n")
	flag.Parse()

	// 如果 env 没有传参，则打印提示信息并退出程序
	if env == nil {
		fmt.Println("请指定运行环境")
		flag.PrintDefaults() // 打印命令行参数的默认值说明
		os.Exit(1)           // 退出程序并返回错误码
	}

	fmt.Println("环境:", *env)

	port := 8081
	if *env == "prod" {
		port = 9091
	}

	engine := gin.Default()
	engine.Use(Recover)

	//初始化路由
	router.InitRouter(engine)

	srv := &http.Server{
		Addr:    ":" + fmt.Sprintf("%v", port),
		Handler: engine,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

//
//func initLog(engine *gin.Engine) {
//	loggersConfig := &logrus.LoggersConfig{
//		LogPath:       config.Log.LogPath,
//		LogLevel:      config.Log.Loglevel,
//		MaxHistory:    time.Duration(config.Log.MaxHistory) * 24 * time.Hour,
//		RollingPolicy: time.Duration(config.Log.RollingPolicy) * 24 * time.Hour,
//		AppName:       configs.GetConfig().App.Info.Name,
//		ConsoleLog:    false,
//		Format:        constant.CommonFormat,
//	}
//	loggersConfig.InitLoggersLogger()
//	loggersConfig.InitGinAccessLog(engine)
//	logrus.Infof("logrus init success,config:%s", mapper.ToJson(loggersConfig))
//}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in f", r)
		}
	}()

	c.Next()
}
