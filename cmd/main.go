package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/DarkSoul94/password-generator/cmd/grpcserver"
	"github.com/DarkSoul94/password-generator/cmd/httpserver"
	"github.com/DarkSoul94/password-generator/pkg/config"
	"github.com/DarkSoul94/password-generator/pkg/logger"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	logger.InitLogger()
	appHttp := httpserver.NewApp()
	go appHttp.Run(viper.GetString("app.http_port"))

	appgRPC := grpcserver.NewApp()
	go appgRPC.Run(viper.GetString("app.grpc_port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	appHttp.Stop()
	appgRPC.Stop()
}
