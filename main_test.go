package appctx_test

import (
	"os"
	"testing"

	"github.com/srvc/appctx"
)

func TestMain(m *testing.M) {
	appctx.ErrorLog = &fakeLogger{
		PrintFunc: func(v ...interface{}) {},
		FatalFunc: func(v ...interface{}) {},
	}
	code := m.Run()
	defer os.Exit(code)
}

type fakeLogger struct {
	PrintFunc func(v ...interface{})
	FatalFunc func(v ...interface{})
}

func (l *fakeLogger) Print(v ...interface{}) { l.PrintFunc(v...) }
func (l *fakeLogger) Fatal(v ...interface{}) { l.FatalFunc(v...) }
