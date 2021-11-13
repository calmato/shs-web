package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/google/uuid"
)

type Client interface {
	Write(ctx context.Context, path string, data []byte) (string, error)
}

type client struct {
	name   string
	bucket *storage.BucketHandle
}

const publicURL = "https://storage.googleapis.com/%s/%s"

// NewClient Cloud Storage接続用のクライアントの生成
func NewClient(ctx context.Context, app *firebase.App, name string) (Client, error) {
	storage, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}
	bucket, err := storage.Bucket(name)
	if err != nil {
		return nil, err
	}
	return &client{name: name, bucket: bucket}, nil
}

func (c *client) Write(ctx context.Context, path string, data []byte) (string, error) {
	filename := fmt.Sprintf("%s/%s", path, uuid.New().String())

	w := c.bucket.Object(filename).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.CacheControl = "public, max-age=86400" // 1 day

	r := bytes.NewReader(data)

	_, err := io.Copy(w, r)
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(publicURL, c.name, filename), nil
}
