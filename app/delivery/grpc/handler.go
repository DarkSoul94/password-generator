package grpc

import (
	"context"

	"github.com/DarkSoul94/password-generator/app"
	pb "github.com/DarkSoul94/password-generator/proto"
)

type Handler struct {
	uc app.Usecase
	pb.UnimplementedPasswordGeneratorServer
}

func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) Generate(ctx context.Context, in *pb.PassParam) (*pb.GenResult, error) {
	pass, err := h.uc.GeneratePassword(int(in.Length), int(in.DigitsCount), in.WithUpper, in.AllowRepeat)
	if err != nil {
		return nil, err
	}

	return &pb.GenResult{Password: pass}, nil
}
