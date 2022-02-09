package backoff

import (
	"math"
	"sync/atomic"
	"time"
)

type Backoff interface {
	Continue() bool
	Wait() <-chan time.Time
	Reset()
}

type fixedInterval struct {
	interval time.Duration
	count    int64
	max      int64
}

type exponent struct {
	count int64
	max   int64
}

// インターバルに応じたリトライ操作を実行用
func NewFixedIntervalBackoff(interval time.Duration, max int64) Backoff {
	return &fixedInterval{
		interval: interval,
		max:      max,
	}
}

func (f *fixedInterval) Continue() bool {
	return f.count <= f.max
}

func (f *fixedInterval) Wait() <-chan time.Time {
	defer func() {
		f.count++
	}()
	if f.interval == 0 {
		c := make(chan time.Time, 1)
		c <- time.Now()
		return c
	}
	return time.NewTimer(f.interval).C
}

func (f *fixedInterval) Reset() {
	f.count = 0
}

// 指数関数のアルゴリズムに応じたリトライ操作を実行用
func NewExponentialBackoff(max int64) Backoff {
	return &exponent{
		max: max,
	}
}

func (e *exponent) Continue() bool {
	return atomic.LoadInt64(&e.count) <= e.max
}

func (e *exponent) Wait() <-chan time.Time {
	defer atomic.AddInt64(&e.count, 1)
	count := atomic.LoadInt64(&e.count)
	wait := time.Second * time.Duration(math.Exp2(float64(count)))
	return time.NewTimer(wait).C
}

func (e *exponent) Reset() {
	atomic.StoreInt64(&e.count, 0)
}
