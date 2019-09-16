package appctx_test

import (
	"context"
	"reflect"
	"syscall"
	"testing"
	"time"

	"github.com/srvc/appctx"
)

func TestWithSignal(t *testing.T) {
	ctx, cancel := appctx.WithSignal(context.Background(), syscall.SIGTERM)
	defer cancel()

	if got, want := ctx.Err(), error(nil); !reflect.DeepEqual(got, want) {
		t.Errorf("context.Err() returns %v, want %v", got, want)
	}

	time.Sleep(5 * time.Millisecond)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	time.Sleep(5 * time.Millisecond)

	if got, want := ctx.Err(), context.Canceled; !reflect.DeepEqual(got, want) {
		t.Errorf("context.Err() returns %v, want %v", got, want)
	}
}
