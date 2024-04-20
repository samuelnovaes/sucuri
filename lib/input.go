package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Input(ctx *context.Context, args ...ast.Expression) ast.Expression {
	var value string
	if len(args) > 0 {
		prompt := args[0].GetValue()
		fmt.Print(prompt)
	}
	_, error := fmt.Scan(&value)
	if error != nil {
		panic(error.Error())
	}
	return &ast.String{Value: value}
}
