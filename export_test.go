package appctx

import "sync"

func ResetGlobal() {
	globalCtxOnce = sync.Once{}
	globalCtx = nil
}
