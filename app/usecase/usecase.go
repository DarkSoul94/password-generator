package usecase

import (
	"github.com/DarkSoul94/password-generator/app"
	"github.com/DarkSoul94/password-generator/pkg/generator"
)

type usecase struct {
	Generator generator.PasswordGenerator
}

func NewUsecase() app.Usecase {
	return &usecase{
		Generator: generator.NewGenerator(),
	}
}

func (u *usecase) GeneratePassword(length, digitsCount int, withUpper, allowRepeat bool) (string, error) {
	return u.Generator.Generate(length, digitsCount, withUpper, allowRepeat)
}
