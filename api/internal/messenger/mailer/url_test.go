package mailer

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeacherURLMaker(t *testing.T) {
	t.Parallel()
	webURL, err := url.Parse("http://example.com")
	require.NoError(t, err)
	maker := NewTeacherURLMaker(webURL)
	res := maker.SignIn()
	assert.Equal(t, "http://example.com/signin", res)
}

func TestStudentURLMaker(t *testing.T) {
	t.Parallel()
	webURL, err := url.Parse("http://example.com")
	require.NoError(t, err)
	maker := NewStudentURLMaker(webURL)
	res := maker.SignIn()
	assert.Equal(t, "http://example.com/signin", res)
}
