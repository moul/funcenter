package main

import (
	"context"
	"log"

	"moul.io/funcenter"
)

func main() {
	ctx, err := funcenter.New(
		context.Background(),
		funcenter.OnInitMiddleware(func(ctx context.Context) (context.Context, error) {
			log.Println("hello from OnInitMiddleware")
			return ctx, nil
		}),
		funcenter.OnEnterMiddleware(func(ctx context.Context, args ...interface{}) context.Context {
			log.Println("hello from OnEnterMiddleware", args)
			return ctx
		}),
		funcenter.OnLeaveMiddleware(func(ctx context.Context) {
			log.Println("hello from OnLeaveMiddleware")
		}),
		funcenter.OnCloseMiddleware(func(ctx context.Context) {
			log.Println("hello from OnCloseMiddleware")
		}),
	)
	if err != nil {
		panic(err)
	}
	defer funcenter.Close(ctx)

	log.Println("hello from main")
	if err := a(ctx, "hello world!"); err != nil {
		log.Println(err)
	}
	log.Println("hello again from main")
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
