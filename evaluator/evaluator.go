package evaluator

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
)

func evalArg(arg ast.Expression, ctx *ast.Context) ast.Expression {
	argKind := arg.GetKind()
	if argKind == ast.IDENTIFIER && !arg.(*ast.Identifier).Evaluate {
		return arg
	}
	return EvalExpression(arg, ctx)
}

func evalCall(call ast.Call, ctx *ast.Context) ast.Expression {
	caller := EvalExpression(call.Caller, ctx)

	if caller.GetKind() == ast.LIB {
		evalArgs := []ast.Expression{}
		for _, arg := range call.Args {
			evalArgs = append(evalArgs, evalArg(arg, ctx))
		}
		fn := *caller.(*ast.Lib).Function
		return fn(ctx, evalArgs...)
	}

	fn := caller.(*ast.Function)
	for i, ident := range fn.Args {
		var argValue ast.Expression = &ast.Null{}
		if len(call.Args) >= i+1 {
			arg := call.Args[i]
			if arg.GetKind() == ast.IDENTIFIER {
				identRef := arg.(*ast.Identifier)
				if identRef.Evaluate {
					argValue = ast.GetRef(identRef, ctx).Value
				} else {
					argValue = (*ctx)[arg.(*ast.Identifier).Symbol]
				}
			} else {
				argValue = arg
			}
		}
		argEvaluated := evalArg(argValue, ctx)
		(*ctx)[ident.Symbol] = &ast.Reference{Value: argEvaluated, Const: true}
	}

	result, _, _ := EvalFunction(*fn, ctx, &ast.Null{})
	return result
}

func evalIdentifier(node ast.Identifier, ctx *ast.Context) ast.Expression {
	ref, ok := (*ctx)[node.Symbol]
	if !ok {
		panic(fmt.Sprintf("%s is not defined", node.Symbol))
	}
	return ref.Value
}

func validateIf(node ast.Expression, ctx *ast.Context) {
	if node.GetKind() != ast.CALL {
		(*ctx)["#IF"] = nil
		return
	}
	caller := node.(*ast.Call).Caller
	if caller.GetKind() != ast.IDENTIFIER {
		(*ctx)["#IF"] = nil
		return
	}
	symbol := caller.(*ast.Identifier).Symbol
	if symbol != "elif" && symbol != "else" {
		(*ctx)["#IF"] = nil
	}
}

func EvalExpression(node ast.Expression, ctx *ast.Context) ast.Expression {
	kind := node.GetKind()
	if kind == ast.CALL {
		return evalCall(*node.(*ast.Call), ctx)
	}
	if kind == ast.IDENTIFIER {
		return evalIdentifier(*node.(*ast.Identifier), ctx)
	}
	return node
}

func EvalFunction(fn ast.Function, ctx *ast.Context, this ast.Expression) (ast.Expression, *ast.Context, bool) {
	ctxCopy := ast.Context{}
	for key, value := range *ctx {
		ctxCopy[key] = value
	}
	ctxCopy["#IF"] = nil
	ctxCopy["this"] = &ast.Reference{Value: this, Const: true}

	for _, expr := range fn.Body {
		validateIf(expr, &ctxCopy)
		exprVal := EvalExpression(expr, &ctxCopy)
		if exprVal.GetKind() == ast.RETURN {
			return exprVal.(*ast.Return).Value, &ctxCopy, true
		}
	}

	return &ast.Null{}, &ctxCopy, false
}
