package backoff

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFixedIntervalBackoff(t *testing.T) {
	t.Parallel()
	var counter int
	backoff := NewFixedIntervalBackoff(time.Second, 2)
	for backoff.Continue() {
		<-backoff.Wait()
		counter++
	}
	assert.Equal(t, 3, counter)
	backoff.Reset()

	backoff = NewFixedIntervalBackoff(0, 2)
	for backoff.Continue() {
		<-backoff.Wait()
		counter++
	}
	assert.Equal(t, 6, counter)
	backoff.Reset()
}

func TestExponentialBackoff(t *testing.T) {
	t.Parallel()
	var counter int
	backoff := NewExponentialBackoff(2)
	for backoff.Continue() {
		<-backoff.Wait()
		counter++
	}
	assert.Equal(t, 3, counter)
	backoff.Reset()
}
