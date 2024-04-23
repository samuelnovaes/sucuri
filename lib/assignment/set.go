package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Set(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	ref := ast.GetRef(args[0], ctx)
	if ref == nil {
		panic("Inavlid SET operation")
	}
	if ref.Const {
		panic("Inavlid SET operation")
	}
	ref.Value = args[1]
	return &ast.Null{}
}
