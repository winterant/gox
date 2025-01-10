package xlog

import "context"

var contextArgsKey int

// ContextWithArgs returns a context with key-values which myslog will print.
func ContextWithArgs(ctx context.Context, kvs ...any) context.Context {
	var args []any
	if ctxKv := ctx.Value(&contextArgsKey); ctxKv != nil {
		args = ctxKv.([]any)
	}
	args = append(args, kvs...)
	return context.WithValue(ctx, &contextArgsKey, args)
}

func getContextArgs(ctx context.Context) []any {
	v := ctx.Value(&contextArgsKey)
	if v != nil {
		return v.([]any)
	}
	return nil
}
