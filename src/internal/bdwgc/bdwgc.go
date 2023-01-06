//go:build gc.bdwgc

package bdwgc

/*
#cgo CFLAGS: -Ilib -Ilib/include -D_WASI_EMULATED_SIGNAL -DHAVE_CONFIG_H -Iwasi

#include <gc/gc.h>

// Declare a C function with the same name as our callback in Go to be able
// to pass it as a function pointer.
void onCollectionEvent();

#include "malloc.c"
*/
import "C"
import "unsafe"

type GCEventType uint32

const (
	GC_EVENT_START GCEventType = C.GC_EVENT_START
)

func GC_malloc(size uintptr) unsafe.Pointer {
	return C.GC_malloc(C.size_t(size))
}

func GC_free(ptr unsafe.Pointer) {
	C.GC_free(ptr)
}

func GC_gcollect() {
	C.GC_gcollect()
}

var onCollectionEventFunc func(eventType GCEventType)

func GC_set_on_collection_event(f func(eventType GCEventType)) {
	onCollectionEventFunc = f
}

//export onCollectionEvent
func onCollectionEvent(eventType uint32) {
	if onCollectionEventFunc != nil {
		onCollectionEventFunc(GCEventType(eventType))
	}
}

func init() {
	C.GC_set_on_collection_event(C.GC_on_collection_event_proc(C.onCollectionEvent))
}
