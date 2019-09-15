# appctx
[![GoDoc](https://godoc.org/github.com/izumin5210/appctx?status.svg)](https://godoc.org/github.com/izumin5210/appctx)
[![GitHub release](https://img.shields.io/github/release/izumin5210/appctx.svg)](https://github.com/izumin5210/appctx/releases/latest)
[![GitHub](https://img.shields.io/github/license/izumin5210/appctx.svg)](./LICENSE)

## Examples
### Global

```
// canceled this context when received os signals for termination
ctx := appctx.Global()
```

### Clone

```go
baseCtx, cancel := context.WithCancel(context.Background())
baseCtx = context.WithValue(baseCtx, "user_id", 123)

newCtx := appctx.Clone(baseCtx)
cancel()

fmt.Println(newCtx.Value("user_id"), baseCtx.Value("user_id")) // => 123 123
fmt.Println(newCtx.Err(), baseCtx.Err()) // => <nil> context canceled
```
