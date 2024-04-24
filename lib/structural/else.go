package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func Else(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if (*ctx)["#IF"] == nil {
		panic("Invalid ELSE operation")
	}
	if !(*ctx)["#IF"].Value.(*ast.Boolean).Value {
		result, _, returned := evaluator.EvalFunction(*args[0].(*ast.Function), ctx)
		if returned {
			return &ast.Return{Value: result}
		}
		if result.GetKind() == ast.BREAK || result.GetKind() == ast.CONTINUE {
			return result
		}
	}
	(*ctx)["#IF"] = nil
	return &ast.Null{}
}
