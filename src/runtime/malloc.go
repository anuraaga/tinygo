//go:build !tinygo.wasm && !custommalloc
// +build !tinygo.wasm,!custommalloc

package runtime

func isMallocPointer(ptr uintptr) bool {
	return false
}
