package appctx_test

import (
	"context"
	"fmt"

	"github.com/izumin5210/appctx"
)

func ExampleClone() {
	debug := func(ctx context.Context) {
		fmt.Println(
			"user_id: ", ctx.Value("user_id"), "\t",
			"request_id: ", ctx.Value("request_id"), "\t",
			"error: ", ctx.Err(),
		)
	}

	baseCtx, cancel := context.WithCancel(context.Background())
	baseCtx = context.WithValue(baseCtx, "user_id", 123)
	baseCtx = context.WithValue(baseCtx, "request_id", "7de4f345-8a68-47bc-a1df-90299f95d753")

	newCtx := appctx.Clone(baseCtx)
	cancel()

	// Output:
	// user_id:  123 	 request_id:  7de4f345-8a68-47bc-a1df-90299f95d753 	 error:  context canceled
	// user_id:  123 	 request_id:  7de4f345-8a68-47bc-a1df-90299f95d753 	 error:  <nil>
	debug(baseCtx)
	debug(newCtx)
}
