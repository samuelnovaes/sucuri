package evaluator

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib"
)

func callFunction(fn ast.Function, ctx ast.Context) ast.Expression {
	for _, expr := range fn.Body {
		exprVal := evalExpression(expr, &ctx)
		if exprVal.GetKind() == ast.RETURN {
			return exprVal.(*ast.Return).Value
		}
	}
	return &ast.Null{}
}

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
	ctxCopy := *ctx
	for i, ident := range fn.Args {
		var argValue ast.Expression = &ast.Null{}
		if len(call.Args) >= i+1 {
			argValue = call.Args[i]
		}
		(*ctx)[ident.Symbol] = evalArg(argValue, &ctxCopy)
	}
	return callFunction(*fn, *ctx)
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

func Eval(program ast.Function) ast.Expression {
	ctx := ast.Context{}
	for key, fn := range lib.Lib {
		ctx[key] = &ast.Lib{Function: &fn}
	}
	return callFunction(program, ctx)
}
