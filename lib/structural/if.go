package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func If(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() != ast.BOOLEAN || args[1].GetKind() != ast.FUNCTION {
		panic("Invalid IF operation")
	}
	condition := args[0].(*ast.Boolean).Value
	(*ctx)["#IF"] = &ast.Reference{Value: &ast.Boolean{Value: condition}, Const: true}
	if condition {
		evaluator.EvalFunction(*args[1].(*ast.Function), ctx)
	}
	return &ast.Null{}
}
