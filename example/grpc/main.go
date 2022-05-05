package main

import (
	"context"
	"fmt"

	pb "github.com/DarkSoul94/password-generator/proto"
	"google.golang.org/grpc"
)

func main() {
	var (
		passLen     int
		digitsCount int
		withUpper   bool
		allowRepeat bool
	)

	fmt.Println("Enter pass len:")
	fmt.Scan(&passLen)

	fmt.Println("Enter digits count:")
	fmt.Scan(&digitsCount)

	fmt.Println("Enter with upper:")
	fmt.Scan(&withUpper)

	fmt.Println("Enter allow repeat:")
	fmt.Scan(&allowRepeat)

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewPasswordGeneratorClient(conn)

	res, err := client.Generate(context.Background(), &pb.PassParam{
		Length:      int32(passLen),
		DigitsCount: int32(digitsCount),
		WithUpper:   withUpper,
		AllowRepeat: allowRepeat,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Password)
}
