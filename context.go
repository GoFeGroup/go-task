package go_task

type Context struct {
	values   map[any]any
	canceled bool
}

func (ctx *Context) IsCanceled() bool {
	return ctx.canceled
}

func (ctx *Context) setCanceled() {
	ctx.canceled = true
}

func (ctx *Context) GetValue(key any) (any, bool) {
	val, ok := ctx.values[key]
	return val, ok
}

func (ctx *Context) SetValue(key, val any) {
	ctx.values[key] = val
}

func innerGetValue[V any](values map[any]any, key any) (rv V, rb bool) {
	val, ok := values[key]
	if !ok {
		return rv, false
	}
	if r, ok := val.(V); ok {
		return r, true
	}
	return rv, false
}

func (ctx *Context) GetString(key any) (string, bool) {
	return innerGetValue[string](ctx.values, key)
}

func (ctx *Context) GetInt(key any) (int, bool) {
	return innerGetValue[int](ctx.values, key)
}

func newContext() *Context {
	return &Context{values: make(map[any]any)}
}
