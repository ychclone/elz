package listener

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	colloc "github.com/elz-lang/elz/collection"
	"github.com/elz-lang/elz/parser"
	_ "llvm.org/llvm/bindings/go/llvm"
)

type ElzListener struct {
	*parser.BaseElzListener
	// AstList contain top level's ast
	AstList []ast.Ast
	// exprStack help we implement expression percedence table.
	exprStack *colloc.Stack // Stack Pop nil is nothing in there
	// exportThis markup the reference Name should be public or not.
	exportThis bool
	// variable default immutable.
	immutable bool
}

func NewElzListener() *ElzListener {
	return &ElzListener{
		immutable: true,
		exprStack: colloc.NewStack(),
	}
}

func (s *ElzListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println(`Elz prog`)
}

func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	fmt.Print(`public `)
	s.exportThis = true
}

func (s *ElzListener) EnterVarDefine(ctx *parser.VarDefineContext) {
	fmt.Print(`var `)
	if ctx.GetMut() != nil {
		s.immutable = false
	}
}

func (s *ElzListener) ExitVarDefine(*parser.VarDefineContext) {
	if !s.immutable {
		s.immutable = true
	}
}

func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	expr := s.exprStack.Pop()
	typ := expr.(ast.Expr).Type()
	fmt.Print(ctx.ID().GetText(), `: `, typ, ` = `)
	fmt.Println(expr)
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}
	s.AstList = append(s.AstList, &ast.VarDefination{
		Immutable:  s.immutable,
		Export:     s.exportThis,
		Name:       ctx.ID().GetText(),
		VarType:    typ,
		Expression: expr.(ast.Expr),
	})
}

func (s *ElzListener) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	le := s.exprStack.Pop()
	re := s.exprStack.Pop()
	e := &ast.BinaryExpr{
		LeftE:  le.(ast.Expr),
		RightE: re.(ast.Expr),
		Op:     ctx.GetOp().GetText(),
	}
	s.exprStack.Push(e)
}

func (s *ElzListener) ExitMulOrDiv(ctx *parser.MulOrDivContext) {
	le := s.exprStack.Pop()
	re := s.exprStack.Pop()
	e := &ast.BinaryExpr{
		LeftE:  le.(ast.Expr),
		RightE: re.(ast.Expr),
		Op:     ctx.GetOp().GetText(),
	}
	s.exprStack.Push(e)
}

func (s *ElzListener) ExitStr(ctx *parser.StrContext) {
	s.exprStack.Push(&ast.Str{Val: ctx.STRING().GetText()})
}
func (s *ElzListener) ExitId(ctx *parser.IdContext) {
	// FIXME: Identifier still do not have Expr to represented it.
	s.exprStack.Push(ctx.ID().GetText())
}
func (s *ElzListener) ExitNum(ctx *parser.NumContext) {
	s.exprStack.Push(&ast.Number{Val: ctx.NUM().GetText()})
}