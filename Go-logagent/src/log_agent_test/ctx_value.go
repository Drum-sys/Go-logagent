package main

import (
	"context"
	"fmt"
)

func main()  {
	ctx := context.WithValue(context.Background(), "trace_id", 123423)
	// 值可以传递
	ctx = context.WithValue(ctx, "session", "ljw")
	process(ctx)
}

func process(ctx context.Context)  {
	ret, ok:= ctx.Value("trace_id").(int)
	if !ok {
		ret = 224
	}
	fmt.Println(ret)

	session := ctx.Value("session").(string)
	fmt.Println(session)
}
