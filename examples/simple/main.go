package main

import (
	"context"
	"fmt"

	"moul.io/funcenter"
)

func main() {
	ctx := funcenter.New(context.Background())
	fmt.Println("hello from main")
	if err := a(ctx, "hello world!"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("hello again from main")
}

func a(ctx context.Context, arg1 string) error {
	ctx = funcenter.EnterFunc(ctx, arg1)
	defer funcenter.ExitFunc(ctx)

	fmt.Println("hello from a")
	return b(ctx, 42, []string{"hello", "world!"})
}

func b(ctx context.Context, arg1 int, arg2 []string) error {
	ctx = funcenter.EnterFunc(ctx, arg1, arg2)
	defer funcenter.ExitFunc(ctx)

	fmt.Println("hello from b")
	return c(ctx)
}

func c(ctx context.Context) error {
	ctx = funcenter.EnterFunc(ctx)
	defer funcenter.ExitFunc(ctx)

	fmt.Println("hello from c")
	return nil
}
