package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type reqeustIDKey struct{}

func ReqeustID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := r.Context()
		ctx = context.WithValue(ctx, reqeustIDKey{}, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func GetRequestID(ctx context.Context) string {
	reqeustID, ok := ctx.Value(reqeustIDKey{}).(string)
	if !ok {
		return reqeustID
	}

	return ""
}
