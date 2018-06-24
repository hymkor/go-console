package console

import (
	"sync"
	"syscall"
)

type CoordT struct {
	x int16
	y int16
}

func (c CoordT) X() int         { return int(c.x) }
func (c CoordT) Y() int         { return int(c.y) }
func (c CoordT) XY() (int, int) { return int(c.x), int(c.y) }

type SmallRectT struct {
	left   int16
	top    int16
	right  int16
	bottom int16
}

func (s SmallRectT) Left() int   { return int(s.left) }
func (s SmallRectT) Top() int    { return int(s.top) }
func (s SmallRectT) Right() int  { return int(s.right) }
func (s SmallRectT) Bottom() int { return int(s.bottom) }






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
