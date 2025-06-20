package interceptors

import (
	"context"
	"merch-store/internal/logger"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
)

func RecoveryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("[interceptor.Recovery] method: %s; error: %v\n%s", info.FullMethod, r, debug.Stack())
			span := opentracing.SpanFromContext(ctx)
			if span == nil {
				return
			}
			ext.Error.Set(span, true)
			span.SetTag("error.message", r)
		}
	}()

	return handler(ctx, req)
}
