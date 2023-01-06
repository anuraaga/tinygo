//go:build (!gc.conservative && !gc.bdwgc) || !tinygo.wasm

package task

type gcData struct{}

func (gcd *gcData) swap() {
}
