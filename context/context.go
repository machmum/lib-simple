package libctx

import "context"

const RequestID = "id"

type ReqIDContextKey string

// NewRequestIDContext returns a new Context carrying requestID.
func NewRequestIDContext(ctx context.Context, requestID string) context.Context {
	k := ReqIDContextKey(RequestID)
	return context.WithValue(ctx, k, requestID)
}

// FromRequestIDContext extracts the user requestID from ctx, if present.
func FromRequestIDContext(ctx context.Context) (value string) {
	if ctx == nil {
		return
	}

	k := ReqIDContextKey(RequestID)
	if v := ctx.Value(k); v != nil {
		value = v.(string)
	}
	return
}
