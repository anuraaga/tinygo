//go:build tinygo.wasm && !custommalloc
// +build tinygo.wasm,!custommalloc

package runtime

import "unsafe"

// The below functions override the default allocator of wasi-libc. This ensures
// code linked from other languages can allocate memory without colliding with
// our GC allocations.

var allocs = make(map[uintptr][]byte)

//export malloc
func libc_malloc(size uintptr) unsafe.Pointer {
	buf := make([]byte, size+4)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	ptr += 4
	allocs[ptr] = buf
	return unsafe.Pointer(ptr)
}

//export free
func libc_free(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	if _, ok := allocs[uintptr(ptr)]; ok {
		delete(allocs, uintptr(ptr))
	} else {
		panic("free: invalid pointer")
	}
}

//export calloc
func libc_calloc(nmemb, size uintptr) unsafe.Pointer {
	// No difference between calloc and malloc.
	return libc_malloc(nmemb * size)
}

//export realloc
func libc_realloc(oldPtr unsafe.Pointer, size uintptr) unsafe.Pointer {
	// It's hard to optimize this to expand the current buffer with our GC, but
	// it is theoretically possible. For now, just always allocate fresh.
	buf := make([]byte, size+4)

	if oldPtr != nil {
		if oldBuf, ok := allocs[uintptr(oldPtr)]; ok {
			copy(buf[4:], oldBuf[4:])
			delete(allocs, uintptr(oldPtr))
		} else {
			panic("realloc: invalid pointer")
		}
	}

	ptr := uintptr(unsafe.Pointer(&buf[0]))
	ptr += 4
	allocs[ptr] = buf
	return unsafe.Pointer(ptr)
}

func isMallocPointer(ptr uintptr) bool {
	_, ok := allocs[ptr]
	return ok
}
