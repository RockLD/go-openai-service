package serve

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-openai-service/conf"
	"go-openai-service/router"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HttpServe struct {
	Logger *zap.Logger
}

func NewHttp() *HttpServe {
	return &HttpServe{
		Logger: conf.NewLogger().InitLogger(viper.GetString("env")),
	}
}

func (http *HttpServe) Run() {
	http.Logger.Info("HttpServe starting...")

	g := gin.New()

	router.InitHttpRouter(g)

	go func() {
		if err := pingServer(); err != nil {
			http.Logger.Error("The router bas been deployed successfully," + err.Error())
		}
	}()

	err := g.Run(":8090")
	if err != nil {
		panic(err)
		//logger.Debug(err)
		http.Logger.Error(err.Error())
		return
	}
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8090" + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
