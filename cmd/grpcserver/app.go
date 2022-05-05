package grpcserver

import (
	"fmt"
	"net"

	appgrpc "github.com/DarkSoul94/password-generator/app/delivery/grpc"
	appusecase "github.com/DarkSoul94/password-generator/app/usecase"
	pb "github.com/DarkSoul94/password-generator/proto"

	"google.golang.org/grpc"
)

type Deps struct {
	PassGenHandler pb.PasswordGeneratorServer
}

type App struct {
	Deps

	grpcServer *grpc.Server
}

func NewApp() *App {
	uc := appusecase.NewUsecase()

	return &App{
		Deps: Deps{
			PassGenHandler: appgrpc.NewHandler(uc),
		},
		grpcServer: grpc.NewServer(),
	}
}

func (a *App) Run(port string) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	pb.RegisterPasswordGeneratorServer(a.grpcServer, a.PassGenHandler)

	a.grpcServer.Serve(l)
}

func (a *App) Stop() {
	a.grpcServer.Stop()
}
