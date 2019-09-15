package appctx

import (
	"context"
	"os"
	"os/signal"
)

// WithSignal returns a copy of parent which is done with os signals.
func WithSignal(parent context.Context, sigs ...os.Signal) (context.Context, context.CancelFunc) {
	return withSignal(parent, func(sig os.Signal) bool { return false }, sigs...)
}

type signalHandlerFunc func(sig os.Signal) bool

func withSignal(parent context.Context, onSignal signalHandlerFunc, sigs ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)

	go func() {
		sigCh := make(chan os.Signal, 1024)
		signal.Notify(sigCh, sigs...)

		defer close(sigCh)
		defer signal.Stop(sigCh)

		for sig := range sigCh {
			cancel()
			if !onSignal(sig) {
				break
			}
		}
	}()

	return ctx, cancel
}
