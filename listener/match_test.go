package listener

import (
	"testing"
	"github.com/dannypsnl/assert"

	"llvm.org/llvm/bindings/go/llvm"
)

func Test_match_rule(t *testing.T) {
	assert := assert.NewTester(t)

	l := listener(`
fn pattern(s: i32) -> i32 {
	match s {
		0 => return 1,
		1 => return s + 1,
		_ => return 0
	}
}`)

	ee, err := llvm.NewExecutionEngine(l.context.Module)
	if err != nil {
		panic(err)
	}
	gv := ee.RunFunction(ee.FindFunction("pattern"), []llvm.GenericValue {
		llvm.NewGenericValueFromInt(llvm.Int32Type() , 0, true),
	})

	assert.Eq(gv.IntWidth(), 32)
	assert.Eq(gv.Int(true), uint64(1))
}
