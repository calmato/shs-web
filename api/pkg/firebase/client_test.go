package firebase

import (
	"context"
	"testing"

	firebase "firebase.google.com/go/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
)

func TestInitializeApp(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name    string
		config  *firebase.Config
		options []option.ClientOption
		isErr   bool
	}{
		{
			name: "success",
			config: &firebase.Config{
				ProjectID: "project-test",
			},
			options: []option.ClientOption{},
			isErr:   false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			app, err := InitializeApp(ctx, tt.config, tt.options...)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, app)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, app)
		})
	}
}
