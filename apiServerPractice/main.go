package main

import (
	"net/http"
	"time"
	"log"
	"errors"

	"github.com/gin-gonic/gin"
	"LearnGolang/apiServerPractice/router"
)

// png server 确保服务启动正常
func pingServer() error {
	for i := 0; i < 2; i++ {
		// ping.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// 休眠一分钟
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}


func main(){
	g := gin.New()
	middleware := []gin.HandlerFunc{}
	//路由加载
	router.Load(g,middleware...)
	go func() {
		if err := pingServer(); err != nil {
		log.Fatal("The router has no response, or it might took too long to start up.", err)
	}
		log.Print("The router has been deployed successfully.")
	}()
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

	}
