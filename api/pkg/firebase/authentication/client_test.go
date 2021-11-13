package authentication

import (
	"context"
	"os"
	"testing"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	setEnv()
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)

	tests := []struct {
		name  string
		app   *firebase.App
		isErr bool
	}{
		{
			name:  "success",
			app:   app,
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			client, err := NewClient(ctx, tt.app)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, client)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, client)
		})
	}
}

func TestAuth_GetUser(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)
	user, err := client.CreateUser(ctx, uid, "test@calmato.jp", "12345678")
	require.NoError(t, err)

	tests := []struct {
		name   string
		uid    string
		expect *auth.UserRecord
		isErr  bool
	}{
		{
			name:   "success",
			uid:    uid,
			expect: user,
			isErr:  false,
		},
		{
			name:   "failed to get user",
			uid:    "",
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			actual, err := client.GetUser(ctx, tt.uid)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestAuth_GetUserByEmail(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)
	user, err := client.CreateUser(ctx, uid, "test@calmato.jp", "12345678")
	require.NoError(t, err)

	tests := []struct {
		name   string
		email  string
		expect *auth.UserRecord
		isErr  bool
	}{
		{
			name:   "success",
			email:  "test@calmato.jp",
			expect: user,
			isErr:  false,
		},
		{
			name:   "failed to get user",
			email:  "",
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			actual, err := client.GetUserByEmail(ctx, tt.email)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestAuth_CreateUser(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)

	tests := []struct {
		name     string
		uid      string
		email    string
		password string
		isErr    bool
	}{
		{
			name:     "success",
			uid:      uid,
			email:    "test@calmato.jp",
			password: "12345678",
			isErr:    false,
		},
		{
			name:     "failed to create user",
			uid:      "",
			email:    "",
			password: "",
			isErr:    true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := client.CreateUser(ctx, tt.uid, tt.email, tt.password)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestAuth_UpdateEmail(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)
	_, err = client.CreateUser(ctx, uid, "test@calmato.jp", "12345678")
	require.NoError(t, err)

	tests := []struct {
		name  string
		uid   string
		email string
		isErr bool
	}{
		{
			name:  "success",
			uid:   uid,
			email: "test-update@calmato.jp",
			isErr: false,
		},
		{
			name:  "failed to update email",
			uid:   "",
			email: "",
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := client.UpdateEmail(ctx, tt.uid, tt.email)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestAuth_UpdatePassword(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)
	_, err = client.CreateUser(ctx, uid, "test@calmato.jp", "12345678")
	require.NoError(t, err)

	tests := []struct {
		name     string
		uid      string
		password string
		isErr    bool
	}{
		{
			name:     "success",
			uid:      uid,
			password: "23456789",
			isErr:    false,
		},
		{
			name:     "failed to update password",
			uid:      "",
			password: "",
			isErr:    true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := client.UpdatePassword(ctx, tt.uid, tt.password)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestAuth_UpdateActivated(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)
	_, err = client.CreateUser(ctx, uid, "test@calmato.jp", "12345678")
	require.NoError(t, err)

	tests := []struct {
		name      string
		uid       string
		activated bool
		isErr     bool
	}{
		{
			name:      "success",
			uid:       uid,
			activated: true,
			isErr:     false,
		},
		{
			name:      "failed to update activated",
			uid:       "",
			activated: false,
			isErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := client.UpdateActivated(ctx, tt.uid, tt.activated)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestAuth_DeleteUser(t *testing.T) {
	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	uid := "00000000-0000-0000-0000-0000000000000"
	_ = client.DeleteUser(ctx, uid)
	_, err = client.CreateUser(ctx, uid, "test@calmato.jp", "12345678")
	require.NoError(t, err)

	tests := []struct {
		name  string
		uid   string
		isErr bool
	}{
		{
			name:  "success",
			uid:   uid,
			isErr: false,
		},
		{
			name:  "failed to delete user",
			uid:   "",
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			err := client.DeleteUser(ctx, tt.uid)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func setEnv() {
	if os.Getenv("FIREBASE_AUTH_EMULATOR_HOST") != "" {
		return
	}
	os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "127.0.0.1:9099")
}
