package webview

/*
#cgo LDFLAGS: -lgdi32

#include <windows.h>

LONG_PTR hIconToLongPtr(HICON hIcon) {
	return (LONG_PTR)hIcon;
}
*/
import "C"

import (
	"image"
	"sync"
	"unsafe"
)

type icons struct {
	mutex sync.Mutex
	icons []C.HICON
}

func (c *icons) setIcon(window unsafe.Pointer, icon image.Image, kind IconKind) {
	hIcon := prepareIcon(icon)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.icons == nil {
		c.icons = make([]C.HICON, maxIconKind)
	}
	previousHIcon := c.icons[kind]
	c.icons[kind] = hIcon
	C.SetClassLongPtr(C.HWND(window), prepareIconKind(kind), C.hIconToLongPtr(hIcon))
	if previousHIcon != nil {
		C.DestroyIcon(previousHIcon)
	}
}

func (c *icons) free() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for i, hIcon := range c.icons {
		if hIcon != nil {
			C.DestroyIcon(hIcon)
			c.icons[i] = nil
		}
	}
}

func prepareIcon(icon image.Image) C.HICON {
	width, height, mask, color := prepareIconData(icon)
	iconInfo := C.ICONINFO{}
	iconInfo.fIcon = C.TRUE
	iconInfo.xHotspot = 0
	iconInfo.yHotspot = 0
	iconInfo.hbmMask = C.CreateBitmap(C.int(width), C.int(height), 1, 8, unsafe.Pointer(&mask[0]))
	iconInfo.hbmColor = C.CreateBitmap(C.int(width), C.int(height), 1, 32, unsafe.Pointer(&color[0]))
	hIcon := C.CreateIconIndirect(&iconInfo)
	C.DeleteObject(C.HGDIOBJ(iconInfo.hbmMask))
	C.DeleteObject(C.HGDIOBJ(iconInfo.hbmColor))
	return hIcon
}

func prepareIconData(img image.Image) (width, height int, mask, color []byte) {
	bounds := img.Bounds()
	width, height = bounds.Dx(), bounds.Dy()
	mask = make([]byte, width*height)
	color = make([]byte, 4*width*height)
	for y, maskIndex, colorIndex := 0, 0, 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			color[colorIndex] = byte(b >> 8)
			colorIndex++
			color[colorIndex] = byte(g >> 8)
			colorIndex++
			color[colorIndex] = byte(r >> 8)
			colorIndex++
			color[colorIndex] = byte(a >> 8)
			colorIndex++
			mask[maskIndex] = byte(a >> 8)
			maskIndex++
		}
	}
	return width, height, mask, color
}

func prepareIconKind(kind IconKind) C.int {
	if kind == IconKindSmaller {
		return C.GCLP_HICONSM
	}
	return C.GCLP_HICON
}
