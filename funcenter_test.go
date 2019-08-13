package funcenter_test

import (
	"context"
	"fmt"

	"moul.io/funcenter"
)

func Example() {
	ctx, err := funcenter.New(
		context.Background(),
		funcenter.OnInitMiddleware(func(ctx context.Context) (context.Context, error) {
			fmt.Println("hello from OnInitMiddleware")
			return ctx, nil
		}),
		funcenter.OnEnterMiddleware(func(ctx context.Context, args ...interface{}) context.Context {
			fmt.Println("hello from OnEnterMiddleware", args)
			return ctx
		}),
		funcenter.OnLeaveMiddleware(func(ctx context.Context) {
			fmt.Println("hello from OnLeaveMiddleware")
		}),
		funcenter.OnCloseMiddleware(func(ctx context.Context) {
			fmt.Println("hello from OnCloseMiddleware")
		}),
	)
	if err != nil {
		panic(err)
	}
	defer funcenter.Close(ctx)

	fmt.Println("hello from main")
	if err := a(ctx, "hello world!"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("hello again from main")
	// Output:
	// hello from OnInitMiddleware
	// hello from main
	// hello from OnEnterMiddleware [hello world!]
	// hello from a
	// hello from OnEnterMiddleware [42 [hello world!]]
	// hello from b
	// hello from OnEnterMiddleware []
	// hello from c
	// hello from OnLeaveMiddleware
	// hello from OnLeaveMiddleware
	// hello from OnLeaveMiddleware
	// hello again from main
	// hello from OnCloseMiddleware
}

func a(ctx context.Context, arg1 string) error {
	ctx = funcenter.Enter(ctx, arg1)
	defer funcenter.Leave(ctx)

	fmt.Println("hello from a")
	return b(ctx, 42, []string{"hello", "world!"})
}

func b(ctx context.Context, arg1 int, arg2 []string) error {
	ctx = funcenter.Enter(ctx, arg1, arg2)
	defer funcenter.Leave(ctx)

	fmt.Println("hello from b")
	return c(ctx)
}

func c(ctx context.Context) error {
	ctx = funcenter.Enter(ctx)
	defer funcenter.Leave(ctx)

	fmt.Println("hello from c")
	return nil
}
