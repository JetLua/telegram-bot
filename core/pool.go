package core

import "sync"

var ctxPool = &sync.Pool{
	New: func() any {
		return &Ctx{}
	},
}
