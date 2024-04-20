package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
)

func Input(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	var value string
	if len(args) > 0 {
		prompt := args[0].GetValue(ctx)
		fmt.Print(prompt)
	}
	_, error := fmt.Scan(&value)
	if error != nil {
		panic(error.Error())
	}
	return &ast.String{Value: value}
}
