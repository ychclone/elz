package ast

import (
	"fmt"
	"strings"
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestArray(t *testing.T) {
	c := NewContext()
	arr := &Array{
		Elements:    []Expr{&I32{Val: "10"}},
		ElementType: "i32",
		Len:         1,
	}

	testCodegenResult(t, arr, c)
}

func testArrayType(t *testing.T, arr *Array, c *Context) {
	actual := arr.Type(c)
	expected := "array<i32,1>"

	if actual != expected {
		t.Errorf(fmt.Sprintf("expected: `%s`, actual: `%s`", expected, actual))
	}
}

func testCodegenResult(t *testing.T, arr *Array, c *Context) {
	arr.Check(c)
	arrv := arr.Codegen(c)
	garr := llvm.AddGlobal(c.Module, llvm.ArrayType(llvm.Int32Type(), 1), "arr")
	garr.SetInitializer(arrv)

	expected := `@arr = global [1 x i32] [i32 10]`

	if !strings.Contains(c.Module.String(), expected) {
		t.Errorf(fmt.Sprintf("expected contains: `%s`, actual module is: `%s`", expected, c.Module))
	}
}
