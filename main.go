package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"go-openai-service/conf"
	"go-openai-service/serve"
	"go.uber.org/zap"
)

var serveType = pflag.StringP("serve", "s", "http", "")
var cfg = pflag.StringP("config", "c", "", "paths")
var ZapLogger *zap.Logger

func init() {
	pflag.Parse()
	if err := conf.InitConfig(*cfg); err != nil { //初始化配置
		fmt.Println("err====", err)
		panic(err)
	}

	ZapLogger = conf.NewLogger().InitLogger("dev")
}

func main() {

	fmt.Println("serveType==========")
	fmt.Println(*serveType)

	switch *serveType {
	case "http":
		serve.NewHttp().Run()
	case "grpc":
		serve.NewGrpc().Run()
	default:
		fmt.Println("Unknown serve type : %s\n", *serveType)
	}

}
