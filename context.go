package go_task

import "sync"

type Context struct {
	//values   map[any]any
	values   sync.Map
	canceled bool
}

func (ctx *Context) IsCanceled() bool {
	return ctx.canceled
}

func (ctx *Context) setCanceled() {
	ctx.canceled = true
}

func (ctx *Context) GetValue(key any) (any, bool) {
	return ctx.values.Load(key)
}

func (ctx *Context) SetValue(key, val any) {
	ctx.values.Store(key, val)
}

func (ctx *Context) GetString(key any) (string, bool) {
	val, ok := ctx.values.Load(key)
	if !ok {
		return "", ok
	}
	if str, ok := val.(string); ok {
		return str, ok
	}
	return "", false
}

func (ctx *Context) GetInt(key any) (int, bool) {
	val, ok := ctx.values.Load(key)
	if !ok {
		return 0, ok
	}
	if n, ok := val.(int); ok {
		return n, ok
	}
	return 0, false
}

func newContext() *Context {
	return &Context{}
}
