//go:build !windows
// +build !windows

package webview

import (
	"image"
	"unsafe"
)

type icons struct{}

func (*icons) setIcon(window unsafe.Pointer, icon image.Image, kind IconKind) {
}

func (*icons) free() {
}
