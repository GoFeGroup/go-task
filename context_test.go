package go_task

import (
	"fmt"
	"testing"
)

func TestContext_Get001(t *testing.T) {
	ctx := newContext()
	fmt.Println(ctx.GetString("3333"))

	ctx.SetValue("3333", "4444")
	fmt.Println(ctx.GetString("3333"))
}

func TestContext_Get002(t *testing.T) {
	ctx := newContext()
	ctx.SetValue("3333", "4444")
	fmt.Println(ctx.GetInt("3333"))
}
