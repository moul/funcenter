package main

import (
	"context"
	"log"

	"go.uber.org/zap"
	"moul.io/funcenter"
	zapmiddleware "moul.io/funcenter/logging/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ctx, _ := funcenter.New(context.Background(), zapmiddleware.New(logger))
	defer funcenter.Close(ctx)

	if err := a(ctx, "hello world!"); err != nil {
		panic(err)
	}
}

func a(ctx context.Context, arg1 string) error {
	ctx = funcenter.Enter(ctx, arg1)
	defer funcenter.Leave(ctx)

	log.Println("hello from a")
	return b(ctx, 42, []string{"hello", "world!"})
}

func b(ctx context.Context, arg1 int, arg2 []string) error {
	ctx = funcenter.Enter(ctx, arg1, arg2)
	defer funcenter.Leave(ctx)

	log.Println("hello from b")
	return c(ctx)
}

func c(ctx context.Context) error {
	ctx = funcenter.Enter(ctx)
	defer funcenter.Leave(ctx)

	log.Println("hello from c")
	return nil
}
