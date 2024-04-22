package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func Else(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if (*ctx)["#IF"] == nil {
		panic("Invalid ELSE operation")
	}
	if !(*ctx)["#IF"].(*ast.Boolean).Value {
		evaluator.EvalFunction(*args[0].(*ast.Function), ctx)
	}
	return &ast.Null{}
}
