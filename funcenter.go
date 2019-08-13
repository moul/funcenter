package funcenter

import (
	"context"
	"fmt"
)

func New(ctx context.Context) context.Context {
	fmt.Println("hello from New")
	return ctx
}

func EnterFunc(ctx context.Context, args ...interface{}) context.Context {
	fmt.Println("hello from EnterFunc")
	return ctx
}

func ExitFunc(ctx context.Context) {
	fmt.Println("hello from ExitFunc")
}
