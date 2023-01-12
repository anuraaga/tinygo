//go:build gc.bdwgc

package runtime

import (
	bdwgc "internal/bdwgc/lib"
	"unsafe"
)

func initHeap() {
	bdwgc.GC_set_on_collection_event(func(eventType bdwgc.GCEventType) {
		switch eventType {
		case bdwgc.GC_EVENT_START:
			markStack()
		}
	})
}

func alloc(size uintptr, layout unsafe.Pointer) unsafe.Pointer {
	buf := bdwgc.GC_malloc(size)
	if buf == nil {
		runtimePanic("out of memory")
	}
	memzero(buf, size)
	return buf
}

// free is called to explicitly free a previously allocated pointer.
func free(ptr unsafe.Pointer) {
	bdwgc.GC_free(ptr)
}

func markRoots(start, end uintptr) {
	// Roots are already registered in bdwgc so we have nothing to do here.
}

func markRoot(start, end uintptr) {
	// Roots are already registered in bdwgc so we have nothing to do here.
}

func GC() {
	bdwgc.GC_gcollect()
}

func setHeapEnd(newHeapEnd uintptr) {
	// Heap is grown by bdwgc, so ignore for when called from wasm initialization.
}
