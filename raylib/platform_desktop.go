// +build !android,!arm

package raylib

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"io"
	"os"
	"unsafe"
)

// Initialize Window and OpenGL Graphics
func InitWindow(width int32, height int32, t interface{}) {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)

	title, ok := t.(string)
	if ok {
		ctitle := C.CString(title)
		defer C.free(unsafe.Pointer(ctitle))
		C.InitWindow(cwidth, cheight, ctitle)
	}
}

// Sets callback function
func SetCallbackFunc(func(unsafe.Pointer)) {
	return
}

// Shows cursor
func ShowCursor() {
	C.ShowCursor()
}

// Hides cursor
func HideCursor() {
	C.HideCursor()
}

// Returns true if cursor is not visible
func IsCursorHidden() bool {
	ret := C.IsCursorHidden()
	v := bool(int(ret) == 1)
	return v
}

// Enables cursor
func EnableCursor() {
	C.EnableCursor()
}

// Disables cursor
func DisableCursor() {
	C.DisableCursor()
}

// Check if a file have been dropped into window
func IsFileDropped() bool {
	ret := C.IsFileDropped()
	v := bool(int(ret) == 1)
	return v
}

// Retrieve dropped files into window
func GetDroppedFiles(count *int32) []string {
	ccount := (*C.int)(unsafe.Pointer(count))
	ret := C.GetDroppedFiles(ccount)

	tmpslice := (*[1 << 24]*C.char)(unsafe.Pointer(ret))[:*count:*count]
	gostrings := make([]string, *count)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}

	return gostrings
}

// Clear dropped files paths buffer
func ClearDroppedFiles() {
	C.ClearDroppedFiles()
}

// Open asset
func OpenAsset(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
