package runtime

import (
	"time"
)

// Assert assert
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

// Retry retry
func Retry(cnt int, fn func() bool) {
	for i := 0; i < cnt; i++ {
		if fn() {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
