package std

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Println(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Print(ctx, append(args, &ast.String{Value: "\n"})...)
}
