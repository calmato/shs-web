package metadata

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

var (
	errInvalidMetadata  = errors.New("metadata: this metadata is invalid")
	errNotFoundMetadata = errors.New("metadata: this metadata is not found")
)

// Get - メタデータの取得
func Get(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errInvalidMetadata
	}

	v := md.Get(key)
	if len(v) == 0 {
		return "", errNotFoundMetadata
	}

	return v[0], nil
}

// Set - メタデータの代入
func Set(ctx context.Context, key string, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
