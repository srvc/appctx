package appctx

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
)

var (
	globalCtx     context.Context
	globalCtxOnce sync.Once

	// Signals are os signals that are handled by Global context in default.
	Signals = []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM}
	// TerminateLimit is a maximum count of received signals until force shutdown.
	TerminateLimit = 3
	// ErrorLog is used for logging received signals.
	ErrorLog Logger = new(defaultLogger)
)

type Logger interface {
	Print(v ...interface{})
	Fatal(v ...interface{})
}

type defaultLogger struct{}

func (defaultLogger) Print(v ...interface{}) { log.Print(v...) }
func (defaultLogger) Fatal(v ...interface{}) { log.Fatal(v...) }

// Global returns a singleton application-scope context that handles termination signals.
func Global() context.Context {
	globalCtxOnce.Do(func() {
		var sigCnt int
		globalCtx, _ = withSignal(context.Background(), func(sig os.Signal) bool {
			sigCnt++
			switch {
			case sigCnt == 1:
				ErrorLog.Print(fmt.Sprintf("Received %v, gracefully stopping", sig))
			case sigCnt >= TerminateLimit:
				ErrorLog.Fatal(fmt.Sprintf("Received signals %d times, Aborting", sigCnt))
				return false
			default:
				ErrorLog.Print(fmt.Sprintf("Received %v", sig))
			}
			return true
		}, Signals...)
	})

	return globalCtx
}
