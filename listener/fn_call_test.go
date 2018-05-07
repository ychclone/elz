package listener

import (
	"testing"

	"strings"
)

func TestFnCall(t *testing.T) {
	r := NewParse(`
	fn add(l, r: i32) -> i32 { return l + r }
	fn main() {
		let a = add(1, 2)
	}
	`)

	expected := `
  %.call_tmp = call i32 @add(i32 1, i32 2)
  %a = alloca i32
  store i32 %.call_tmp, i32* %a`

	if !strings.Contains(r, expected) {
		t.Errorf("expected: %s, actual: %s", expected, r)
	}
}
