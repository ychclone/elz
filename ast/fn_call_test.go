package ast

import (
	"testing"

	"strings"
)

func TestFnCall(t *testing.T) {
	context := NewContext()
	f := &FnDef{
		Export: false,
		Name:   "add",
		Params: []*Param{
			&Param{
				Name: "lv",
				Type: "i32",
			},
			&Param{
				Name: "rv",
				Type: "i32",
			},
		},
		Body: []Stat{
			&Return{
				Expr: &BinaryExpr{
					LeftE:  &Id{Val: "lv"},
					RightE: &Id{Val: "rv"},
					Op:     "+",
				},
			},
		},
		RetType: "i32",
	}
	fc := &FnCall{
		Name: "add",
		Args: []Expr{
			&I32{Val: "10"},
			&I32{Val: "10"},
		},
	}
	m := &FnDef{
		Export: false,
		Name:   "main",
		Params: []*Param{},
		Body: []Stat{
			fc,
		},
		RetType: "i32",
	}
	f.Check(context)
	m.Check(context)

	f.Codegen(context)
	m.Codegen(context)

	expected := `define i32 @add(i32 %lv, i32 %rv) {
entry:
  %.add_tmp = add i32 %lv, %rv
  ret i32 %.add_tmp
}

define i32 @main() {
entry:
  %.call_tmp = call i32 @add(i32 10, i32 10)
  ret i32 0
}`
	if !strings.Contains(context.Module.String(), expected) {
		t.Error()
	}
}
