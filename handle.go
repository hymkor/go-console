package console

import (
	"sync"
	"syscall"
)

// Handle is the alias of syscall.Handle
type Handle = syscall.Handle

var Kernel32 = syscall.NewLazyDLL("kernel32")

var out Handle
var outOnce sync.Once

// ConOut returns the handle for Console-Output
func Out() Handle {
	outOnce.Do(func() {
		var err error
		out, err = syscall.Open("CONOUT$", syscall.O_RDWR, 0)
		if err != nil {
			panic(err.Error())
		}
	})
	return out
}

var in Handle
var inOnce sync.Once

func In() Handle {
	inOnce.Do(func() {
		var err error
		in, err = syscall.Open("CONIN$", syscall.O_RDWR, 0)
		if err != nil {
			panic(err.Error())
		}
	})
	return in
}
