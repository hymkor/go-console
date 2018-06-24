// +build windows

package csbi

import (
	"github.com/zetamatta/go-console"
	"unsafe"
)

type coordT = console.CoordT
type smallRectT = console.SmallRectT

// ConsoleScreenBufferInfoT is the type for structure contains terminal's information.
type ConsoleScreenBufferInfoT struct {
	Size              coordT
	CursorPosition    coordT
	Attributes        uint16
	Window            smallRectT
	MaximumWindowSize coordT
}

var getConsoleScreenBufferInfo = console.Kernel32.NewProc("GetConsoleScreenBufferInfo")

// GetConsoleScreenBufferInfo returns the latest ConsoleScreenBufferInfoT
// cursor position, window region.
func GetConsoleScreenBufferInfo() *ConsoleScreenBufferInfoT {
	var csbi ConsoleScreenBufferInfoT
	getConsoleScreenBufferInfo.Call(
		uintptr(console.Out()),
		uintptr(unsafe.Pointer(&csbi)))
	return &csbi
}

// ViewSize returns window size from ConsoleScreenBufferInfo structure.
func (csbi *ConsoleScreenBufferInfoT) ViewSize() (int, int) {
	return csbi.Window.Right() - csbi.Window.Left() + 1,
		csbi.Window.Bottom() - csbi.Window.Top() + 1
}
