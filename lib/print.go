package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
)

func Print(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	strs := []any{}
	for _, arg := range args {
		strs = append(strs, arg.GetValue(ctx))
	}
	_, error := fmt.Print(strs...)
	if error != nil {
		panic(error.Error())
	}
	return &ast.Null{}
}
