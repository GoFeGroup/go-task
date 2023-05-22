package go_task

import (
	"fmt"
	"testing"
)

func TestContext_Get001(t *testing.T) {
	ctx := NewContext()
	fmt.Println(ctx.GetString("3333"))

	ctx.SetValue("3333", "4444")
	fmt.Println(ctx.GetString("3333"))
}

func TestContext_Get002(t *testing.T) {
	ctx := NewContext()
	ctx.SetValue("3333", "4444")
	fmt.Println(ctx.GetInt("3333"))
}

func TestContext_Get003(t *testing.T) {
	ctx := NewContext()
	ctx.SetValue("3333", 4444)
	fmt.Println(ctx.GetInt("3333"))
}

func TestContext_Get004(t *testing.T) {
	ctx := NewContext()
	ctx.SetValue("3333", 4444)
	fmt.Println(ctx.GetString("3333"))
}
