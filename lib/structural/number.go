package structural

import (
	"strconv"

	"github.com/samuelnovaes/sucuri/ast"
)

func Number(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() != ast.STRING {
		panic("Invalid NUMBER operation")
	}
	originalValue := args[0].(*ast.String).Value
	convertedValue, error := strconv.ParseFloat(originalValue, 64)
	if error != nil {
		panic(error.Error())
	}
	return &ast.Number{Value: convertedValue}
}
