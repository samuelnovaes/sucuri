package evaluator

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func evalArg(arg ast.Expression, ctx *ast.Context) ast.Expression {
	argKind := arg.GetKind()
	if argKind == ast.IDENTIFIER && !arg.(*ast.Identifier).Evaluate {
		return arg
	}
	return evalExpression(arg, ctx)
}

func evalCall(call ast.Call, ctx *ast.Context) ast.Expression {
	caller := evalExpression(call.Caller, ctx)

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
			argValue = call.Args[i]
		}
		(*ctx)[ident.Symbol] = evalArg(argValue, ctx)
	}

	return EvalFunction(*fn, ctx)
}

func evalIdentifier(node ast.Identifier, ctx *ast.Context) ast.Expression {
	val := (*ctx)[node.Symbol]
	if val == nil {
		val = &ast.Null{}
	}
	return val
}

func evalExpression(node ast.Expression, ctx *ast.Context) ast.Expression {
	kind := node.GetKind()
	if kind == ast.CALL {
		return evalCall(*node.(*ast.Call), ctx)
	}
	if kind == ast.IDENTIFIER {
		return evalIdentifier(*node.(*ast.Identifier), ctx)
	}
	return node
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

func EvalFunction(fn ast.Function, ctx *ast.Context) ast.Expression {
	ctxCopy := ast.Context{}
	for key, value := range *ctx {
		ctxCopy[key] = value
	}
	ctxCopy["#IF"] = nil

	for _, expr := range fn.Body {
		validateIf(expr, &ctxCopy)
		exprVal := evalExpression(expr, &ctxCopy)
		if exprVal.GetKind() == ast.RETURN {
			return exprVal.(*ast.Return).Value
		}
	}

	return &ast.Null{}
}
