package recovery

import "github.com/lwch/logging"

// Recover recover
func Recovery() {
	if err := recover(); err != nil {
		logging.Error("panic: %v", err)
	}
}
