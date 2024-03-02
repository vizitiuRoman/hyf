package main

import (
	"context"
	"fmt"

	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	hyfv1 "github.com/vizitiuRoman/hyf/pkg/adapter/hyf/v1"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
)

func main() {
	logger := log.MustDefaultConsoleLogger("debug")

	client, err := hyfv1.NewTodoSVCClient(
		context.Background(),
		logger,
		&hyfv1.Config{
			Host:     "localhost",
			GrpcPort: 32772,
		},
		nil, nil,
	)

	out, err := client.CreateTodo(context.Background(), &pb.CreateTodoIn{
		Todo: &pb.Todo{
			Name:        "GG",
			Description: "QQ",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out.String())
}
