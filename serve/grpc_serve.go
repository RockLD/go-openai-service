package serve

import (
	"github.com/spf13/viper"
	"go-openai-service/conf"
	"go.uber.org/zap"
)

type GrpcServe struct {
	Logger *zap.Logger
}

func NewGrpc() *GrpcServe {
	return &GrpcServe{
		Logger: conf.NewLogger().InitLogger(viper.GetString("env")),
	}
}

func (g *GrpcServe) Run() {
	g.Logger.Info("GrpcServe starting...")
}
