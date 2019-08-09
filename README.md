# clonectx
[![GoDoc](https://godoc.org/github.com/izumin5210/clonectx?status.svg)](https://godoc.org/github.com/izumin5210/clonectx)
[![GitHub release](https://img.shields.io/github/release/izumin5210/clonectx.svg)](https://github.com/izumin5210/clonectx/releases/latest)
[![GitHub](https://img.shields.io/github/license/izumin5210/clonectx.svg)](./LICENSE)

```go
baseCtx, cancel := context.WithCancel(context.Background())
baseCtx = context.WithValue(baseCtx, "user_id", 123)

newCtx := clonectx.Clone(baseCtx)
cancel()

fmt.Println(newCtx.Value("user_id"), baseCtx.Value("user_id")) // => 123 123
fmt.Println(newCtx.Err(), baseCtx.Err()) // => <nil> context canceled
```
