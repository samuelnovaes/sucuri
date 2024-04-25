package structural

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

type Object map[any]ast.Expression

func getKey(arg ast.Expression, ctx *ast.Context) any {
	if arg.GetKind() == ast.IDENTIFIER {
		return arg.(*ast.Identifier).Symbol
	}
	if arg.GetKind() == ast.STRING || arg.GetKind() == ast.NUMBER {
		return arg.GetValue(ctx)
	}
	panic(fmt.Sprintf("Invalid key \"%s\"", arg.GetValue(ctx)))
}

func getHandler(obj Object) ast.LibFunction {
	return func(ctx *ast.Context, args ...ast.Expression) ast.Expression {
		key := getKey(args[0], ctx)
		if len(args) > 1 {
			obj[key] = args[1]
			return &ast.Null{}
		}
		result, ok := obj[key]
		if ok {
			if result.GetKind() == ast.FUNCTION {
				fn := *result.(*ast.Function)
				var wrapper ast.LibFunction = func(ctx *ast.Context, args ...ast.Expression) ast.Expression {
					this := getHandler(obj)
					result, _, _ := evaluator.EvalFunction(fn, ctx, &ast.Lib{Function: &this})
					return result
				}
				return &ast.Lib{Function: &wrapper}
			}
			return result
		}
		return &ast.Null{}
	}
}

func Obj(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	obj := map[any]ast.Expression{}

	for i, arg := range args {
		if i%2 == 1 {
			obj[getKey(args[i-1], ctx)] = arg
		}
	}

	this := getHandler(obj)
	return &ast.Lib{Function: &this}
}
