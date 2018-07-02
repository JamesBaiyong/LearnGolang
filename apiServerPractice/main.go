package main

import (
	"errors"
	"net/http"
	"time"

	"LearnGolang/apiServerPractice/config"
	"LearnGolang/apiServerPractice/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiServerPractice config file path")
)

// png server 确保服务启动正常
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// ping.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// 休眠一分钟
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}


	gin.SetMode(viper.GetString("runmode"))

	// 测试日志打印并保存
	//for{
	//	log.Info("test for log writer....................................................................................")
	//	time.Sleep(100*time.Millisecond)
	//}

	g := gin.New()
	middleware := []gin.HandlerFunc{}
	//路由加载
	router.Load(g, middleware...)
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}
