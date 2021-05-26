package system

import (
	"fmt"
	"github.com/topfreegames/pitaya/logger"
	"runtime"
	"runtime/debug"


)

// Function is to get current func name with package name
func Function() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// FileLine is to get caller's file name and line number
func FileLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic(fmt.Errorf("get file & line failed"))
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// SeriousPanic is to cause a fatal error when panic ocurrs.
func SeriousPanic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Errorf("panic: %v\n%s", err.(error), string(debug.Stack()))

		}
	}()

	fn()
}

// SlightPanic is to cause no error when panic ocurrs.
func SlightPanic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Errorf("panic: %v\n%s", err.(error), string(debug.Stack()))
		}
	}()

	fn()
}

// SilentPanic is to cause a normal error when panic ocurrs.
func SilentPanic(fn func() error) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("panic: %v\n%s", err.(error), string(debug.Stack()))
		}
	}()

	return fn()
}
