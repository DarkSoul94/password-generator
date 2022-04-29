package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DarkSoul94/password-generator/app"
	apphttp "github.com/DarkSoul94/password-generator/app/delivery/http"
	appusecase "github.com/DarkSoul94/password-generator/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// App ...
type App struct {
	appUC      app.Usecase
	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	uc := appusecase.NewUsecase()
	return &App{
		appUC: uc,
	}
}

// Run run application
func (a *App) Run(port string) error {
	router := gin.New()
	if viper.GetBool("app.release") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}

	apiRouter := router.Group("/api")
	apphttp.RegisterHTTPEndpoints(apiRouter, a.appUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var l net.Listener
	var err error
	l, err = net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		panic(err)
	}

	go func(l net.Listener) {
		if err := a.httpServer.Serve(l); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}(l)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
