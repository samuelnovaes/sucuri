package evaluator

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
	"github.com/samuelnovaes/sucuri/lib"
)

func callFunction(fn ast.Function, ctx context.Context) ast.Expression {
	for _, expr := range fn.Body {
		exprVal := evalExpression(expr, &ctx)
		if exprVal.GetKind() == ast.RETURN {
			return exprVal.(*ast.Return).Value
		}
	}
	return &ast.Null{}
}

func evalArg(arg ast.Expression, ctx *context.Context) ast.Expression {
	argKind := arg.GetKind()
	if (argKind == ast.IDENTIFIER && !arg.(*ast.Identifier).Evaluate) || argKind == ast.LIB {
		return arg
	}
	return evalExpression(arg, ctx)
}

func evalCall(call ast.Call, ctx *context.Context) ast.Expression {
	caller := evalExpression(call.Caller, ctx)
	if caller.GetKind() == ast.LIB {
		libExpr := caller.(*ast.Lib)
		libFn := lib.Lib[libExpr.Symbol]
		evalArgs := []ast.Expression{}
		for _, arg := range call.Args {
			evalArgs = append(evalArgs, evalArg(arg, ctx))
		}
		return libFn(ctx, evalArgs...)
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

func evalIdentifier(node ast.Identifier, ctx *context.Context) ast.Expression {
	val := (*ctx)[node.Symbol]
	if val == nil && lib.Lib[node.Symbol] != nil {
		val = &ast.Lib{Symbol: node.Symbol}
	}
	if val == nil {
		val = &ast.Null{}
	}
	return val
}

func evalExpression(node ast.Expression, ctx *context.Context) ast.Expression {
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
	ctx := context.Context{}
	return callFunction(program, ctx)
}
