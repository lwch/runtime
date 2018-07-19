package aux

import (
	"fmt"
	"runtime"
	"strings"
)

// Exception exception
type Exception []string

func (e Exception) Error() string {
	return strings.Join(e, "\n")
}

func throw(msg string, args ...interface{}) {
	panic(Trace(msg, args...))
}

// Trace stack trace
func Trace(msg string, args ...interface{}) (logs Exception) {
	logs = Exception{fmt.Sprintf(msg, args...)}
	n := 1
	for {
		n++
		pc, file, line, ok := runtime.Caller(n)
		if !ok {
			break
		}
		f := runtime.FuncForPC(pc)
		name := f.Name()
		if strings.HasPrefix(name, "runtime.") {
			continue
		}
		fn := file[strings.Index(file, "/src/")+5:]
		logs = append(logs, fmt.Sprintf("\t(%s:%d) %s", fn, line, name))
	}
	return
}

func catch(err *error, handler ...func()) {
	if e := recover(); e != nil {
		*err = e.(error)
	}
	for _, h := range handler {
		h()
	}
}
