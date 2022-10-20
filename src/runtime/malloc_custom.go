//go:build custommalloc
// +build custommalloc

package runtime

// If custom malloc is wired up, this must be defined using go:linkname
func isMallocPointer(ptr uintptr) bool
