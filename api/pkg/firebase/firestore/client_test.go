package firestore

import (
	"context"
	"os"
	"testing"

	firebase "firebase.google.com/go/v4"
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
			err = client.Close()
			assert.NoError(t, err)
		})
	}
}

func TestFirestore_Get(t *testing.T) {
	type hoge struct {
		Value string `firestore:"value"`
	}

	setEnv()
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	_ = client.DeleteDoc(ctx, "collection", "document")
	err = client.Set(ctx, "collection", "document", &hoge{Value: "debug"})
	require.NoError(t, err)

	tests := []struct {
		name       string
		collection string
		document   string
		expect     *hoge
		isErr      bool
	}{
		{
			name:       "success",
			collection: "collection",
			document:   "document",
			expect:     &hoge{Value: "debug"},
			isErr:      false,
		},
		{
			name:       "not found",
			collection: "collection",
			document:   "",
			expect:     nil,
			isErr:      true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			doc, err := client.Get(ctx, tt.collection, tt.document)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, doc)
				return
			}
			require.NoError(t, err)
			actual := &hoge{}
			assert.NoError(t, doc.DataTo(actual))
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestFirestore_Set(t *testing.T) {
	type hoge struct {
		Value string `firestore:"value"`
	}

	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	_ = client.DeleteDoc(ctx, "collection", "document")

	tests := []struct {
		name       string
		collection string
		document   string
		data       *hoge
		isErr      bool
	}{
		{
			name:       "success",
			collection: "collection",
			document:   "document",
			data:       &hoge{Value: "debug"},
			isErr:      false,
		},
		{
			name:       "invalid argument",
			collection: "",
			document:   "",
			data:       nil,
			isErr:      true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := client.Set(ctx, tt.collection, tt.document, tt.data)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestFirestore_DeleteDoc(t *testing.T) {
	type hoge struct {
		Value string `firestore:"value"`
	}

	setEnv()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &firebase.Config{ProjectID: "project-test"}
	app, err := firebase.NewApp(ctx, config)
	require.NoError(t, err)
	client, err := NewClient(ctx, app)
	require.NoError(t, err)

	_ = client.DeleteDoc(ctx, "collection", "document")
	err = client.Set(ctx, "collection", "document", &hoge{Value: "debug"})
	require.NoError(t, err)

	tests := []struct {
		name       string
		collection string
		document   string
		isErr      bool
	}{
		{
			name:       "success",
			collection: "collection",
			document:   "document",
			isErr:      false,
		},
		{
			name:       "invalid argument",
			collection: "",
			document:   "",
			isErr:      true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := client.DeleteDoc(ctx, tt.collection, tt.document)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func setEnv() {
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		os.Setenv("FIRESTORE_EMULATOR_HOST", "127.0.0.1:8080")
	}
	if os.Getenv("GCP_SERVICE_KEY_JSON") == "" {
		os.Setenv("GCP_SERVICE_KEY_JSON", "{}")
	}
}
