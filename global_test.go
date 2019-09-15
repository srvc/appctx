package appctx_test

import (
	"reflect"
	"syscall"
	"testing"
	"time"

	"github.com/izumin5210/appctx"
)

func TestGlobal(t *testing.T) {
	var printCnt, fatalCnt int
	logger := &fakeLogger{
		PrintFunc: func(v ...interface{}) { printCnt++ },
		FatalFunc: func(v ...interface{}) { fatalCnt++ },
	}
	defer func(l appctx.Logger) { appctx.ErrorLog = l }(appctx.ErrorLog)
	appctx.ErrorLog = logger

	ctx := appctx.Global()
	defer appctx.ResetGlobal()

	if got, want := ctx, appctx.Global(); got != want {
		t.Error("appctx.Global() should return the same object anytime")
	}

	if got, want := ctx.Err(), error(nil); !reflect.DeepEqual(got, want) {
		t.Errorf("ctx.Err() returns %v, want %v", got, want)
	}

	time.Sleep(5 * time.Millisecond)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	time.Sleep(5 * time.Millisecond)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	time.Sleep(5 * time.Millisecond)

	if got, want := printCnt, 2; got != want {
		t.Errorf("ErrorLog.Print called %d times, want %d", got, want)
	}

	if got, want := fatalCnt, 0; got != want {
		t.Errorf("ErrorLog.Fatal called %d times, want %d", got, want)
	}

	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	time.Sleep(5 * time.Millisecond)

	if got, want := printCnt, 2; got != want {
		t.Errorf("ErrorLog.Print called %d times, want %d", got, want)
	}

	if got, want := fatalCnt, 1; got != want {
		t.Errorf("ErrorLog.Fatal called %d times, want %d", got, want)
	}
}
