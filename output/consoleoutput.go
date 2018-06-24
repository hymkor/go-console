package consoleoutput

import (
	"bytes"
	"strings"
	"unsafe"

	"github.com/zetamatta/go-console"
	"github.com/zetamatta/go-console/screenbuffer"
)

type Coord = console.CoordT
type SmallRect = console.SmallRectT

type CharInfoT struct {
	UnicodeChar uint16
	Attributes  uint16
}

const (
	COMMON_LVB_LEADING_BYTE  = 0x0100
	COMMON_LVB_TRAILING_BYTE = 0x0200
)

var readConsoleOutput = console.Kernel32.NewProc("ReadConsoleOutputW")

func ReadConsoleOutput(buffer []CharInfoT, size Coord, coord Coord, read_region *SmallRect) error {

	sizeValue := *(*uintptr)(unsafe.Pointer(&size))
	coordValue := *(*uintptr)(unsafe.Pointer(&coord))

	status, _, err := readConsoleOutput.Call(
		uintptr(console.Out()),
		uintptr(unsafe.Pointer(&buffer[0])),
		sizeValue,
		coordValue,
		uintptr(unsafe.Pointer(read_region)))
	if status == 0 {
		return err
	}
	return nil
}

func GetRecentOutput() (string, error) {
	screen := csbi.GetConsoleScreenBufferInfo()

	y := 0
	h := 1
	if screen.CursorY() >= 1 {
		y = screen.CursorY() - 1
		h++
	}

	region := console.LeftTopRightBottom(0, y, screen.Size.X(), y+h-1)

	home := &Coord{}
	charinfo := make([]CharInfoT, screen.Width()*screen.Height())
	err := ReadConsoleOutput(charinfo, screen.Size, *home, region)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for i := 0; i < screen.Height(); i++ {
		for j := 0; j < screen.Width(); j++ {
			p := &charinfo[i*screen.Width()+j]
			if (p.Attributes & COMMON_LVB_TRAILING_BYTE) != 0 {
				// right side of wide charactor

			} else if (p.Attributes & COMMON_LVB_LEADING_BYTE) != 0 {
				// left side of wide charactor
				if p.UnicodeChar != 0 {
					buffer.WriteRune(rune(p.UnicodeChar))
				}
			} else {
				// narrow charactor
				if p.UnicodeChar != 0 {
					buffer.WriteRune(rune(p.UnicodeChar & 0xFF))
				}
			}
		}
	}
	return strings.TrimSpace(buffer.String()), nil
}
