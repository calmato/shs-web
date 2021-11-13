package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

type Client interface {
	Close() error
	Get(ctx context.Context, collection, document string) (*firestore.DocumentSnapshot, error)
	Set(ctx context.Context, collection, document string, data interface{}) error
	DeleteDoc(ctx context.Context, collection, document string) error
}

type client struct {
	firestore *firestore.Client
}

// NewClient Firestore接続用クライアントの生成
func NewClient(ctx context.Context, app *firebase.App) (Client, error) {
	fs, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &client{firestore: fs}, nil
}

// Close Firestoreとの接続を終了
func (c *client) Close() error {
	return c.firestore.Close()
}

// Get 単一のドキュメントの内容を取得
func (c *client) Get(ctx context.Context, collection, document string) (*firestore.DocumentSnapshot, error) {
	doc, err := c.firestore.Collection(collection).Doc(document).Get(ctx)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Set 単一のドキュメントを作成または上書き
func (c *client) Set(ctx context.Context, collection, document string, data interface{}) error {
	_, err := c.firestore.Collection(collection).Doc(document).Set(ctx, data)
	return err
}

// DeleteDoc ドキュメントを削除
func (c *client) DeleteDoc(ctx context.Context, collection, document string) error {
	_, err := c.firestore.Collection(collection).Doc(document).Delete(ctx)
	return err
}
