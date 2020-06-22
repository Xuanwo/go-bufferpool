package bufferpool

import (
	"testing"
)

func BenchmarkPool_Get(b *testing.B) {
	p := New(1024)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := p.Get()
			x.Free()
		}
	})
}
