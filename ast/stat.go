package ast

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
)

type Stat interface {
	Codegen(*Context) llvm.Value
	Check(*Context)
}

type GlobalVarDef struct {
	Export bool
	// a = 1, a is Name
	Name string
	// a: num = 1, num is VarType, but expression could not have the same type, we have to check it.
	VarType    string
	Expression Expr
}

func (g *GlobalVarDef) Check(ctx *Context) {
	if g.VarType == "" {
		g.VarType = g.Expression.Type(ctx)
	}
	exprType := g.Expression.Type(ctx)
	if g.VarType != exprType {
		ctx.Reporter.Emit(fmt.Sprintf("global var: %s, it's type is: %s, but receive: %s", g.Name, g.VarType, exprType))
	}
	ctx.VarsType[g.Name] = g.VarType
}

func (varDef *GlobalVarDef) Codegen(ctx *Context) llvm.Value {
	// Parser should insert Type if user didn't define it.
	// So we should not get null string
	expr := varDef.Expression.Codegen(ctx)
	val := llvm.AddGlobal(ctx.Module, expr.Type(), varDef.Name)
	val.SetInitializer(expr)
	ctx.Vars[varDef.Name] = expr // Let's use load to access all var
	return val
}

type LocalVarDef struct {
	Immutable  bool
	Name       string
	VarType    string
	Expression Expr
}

func (lv *LocalVarDef) Check(ctx *Context) {
	if lv.VarType == "" {
		lv.VarType = lv.Expression.Type(ctx)
	}
	exprType := lv.Expression.Type(ctx)
	if lv.VarType != exprType {
		ctx.Reporter.Emit(fmt.Sprintf("var: %s, it's type is: %s, but receive: %s", lv.Name, lv.VarType, exprType))
	}
	ctx.VarsType[lv.Name] = lv.VarType
}

func (lv *LocalVarDef) Codegen(ctx *Context) llvm.Value {
	exprType := lv.Expression.Type(ctx)
	if lv.VarType != "" && exprType != lv.VarType {
		ctx.Reporter.Emit(fmt.Sprintf("local var: %s, it's type is: %s, but receive: %s", lv.Name, lv.VarType, exprType))
		return llvm.Value{}
	} else {
		expr := lv.Expression.Codegen(ctx)
		val := ctx.Builder.CreateAlloca(convertToLLVMType(lv.VarType), lv.Name)
		ctx.Builder.CreateStore(expr, val)
		return val
	}
}
