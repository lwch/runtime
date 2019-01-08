package runtime

import (
	"bytes"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// UUID generate uuid string with n
func UUID(n int) (string, error) {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	var ret bytes.Buffer
	for n > 0 {
		buf := make([]byte, n)
		readen, err := rand.Read(buf)
		if err != nil {
			return "", err
		}
		for _, ch := range buf[:readen] {
			if err := ret.WriteByte(chars[int(ch)%len(chars)]); err != nil {
				return "", err
			}
		}
	}
	return ret.String(), nil
}
