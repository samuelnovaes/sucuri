package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Continue(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return &ast.Return{Value: &ast.Continue{}}
}
