package authentication

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type Client interface {
	VerifyIDToken(ctx context.Context, idToken string) (string, error)
	GetUser(ctx context.Context, uid string) (*auth.UserRecord, error)
	GetUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error)
	CreateUser(ctx context.Context, uid string, email string, password string) (*auth.UserRecord, error)
	UpdateEmail(ctx context.Context, uid string, email string) (*auth.UserRecord, error)
	UpdatePassword(ctx context.Context, uid string, password string) (*auth.UserRecord, error)
	UpdateActivated(ctx context.Context, uid string, activated bool) (*auth.UserRecord, error)
	DeleteUser(ctx context.Context, uid string) error
}

type client struct {
	auth *auth.Client
}

// NewClient - Firebase Authentication接続用クライアントの生成
func NewClient(ctx context.Context, app *firebase.App) (Client, error) {
	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &client{auth: auth}, nil
}

// VerifyIDToken IDトークンを検証して、デコードされたトークンから認証ユーザーのUIDを取得
func (c *client) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	token, err := c.auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	return token.UID, nil
}

// GetUser UIDによるユーザー情報の取得
func (c *client) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	return c.auth.GetUser(ctx, uid)
}

// GetUIDByEmail メールアドレスによるユーザー情報の取得
func (c *client) GetUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error) {
	return c.auth.GetUserByEmail(ctx, email)
}

// CreateUser 新しいFirebase Authenticationユーザーを作成
func (c *client) CreateUser(ctx context.Context, uid string, email string, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)
	return c.auth.CreateUser(ctx, params)
}

// UpdateEmail メールアドレスの更新
func (c *client) UpdateEmail(ctx context.Context, uid string, email string) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Email(email).
		EmailVerified(false)
	return c.auth.UpdateUser(ctx, uid, params)
}

// UpdatePassword パスワードの更新
func (c *client) UpdatePassword(ctx context.Context, uid string, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Password(password)
	return c.auth.UpdateUser(ctx, uid, params)
}

// UpdateActivated アカウント状態の変更
func (c *client) UpdateActivated(ctx context.Context, uid string, activated bool) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Disabled(!activated)
	return c.auth.UpdateUser(ctx, uid, params)
}

// DeleteUser UIDによるユーザーの退会
func (c *client) DeleteUser(ctx context.Context, uid string) error {
	return c.auth.DeleteUser(ctx, uid)
}
