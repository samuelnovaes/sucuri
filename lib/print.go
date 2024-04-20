package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Print(ctx *context.Context, args ...ast.Expression) ast.Expression {
	strs := []any{}
	for _, arg := range args {
		strs = append(strs, arg.GetValue())
	}
	_, error := fmt.Print(strs...)
	if error != nil {
		panic(error.Error())
	}
	return &ast.Null{}
}
