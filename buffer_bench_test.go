package bufferpool

import (
	"testing"
	"time"
)

func benchmarkBuffer(b *testing.B, fn func(buf *Buffer)) {
	p := New(1024)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		buf := p.Get()

		for pb.Next() {
			fn(buf)
			buf.Reset()
		}
	})
}

func BenchmarkBuffer_AppendBool(b *testing.B) {
	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendBool(true)
	})
}

func BenchmarkBuffer_AppendByte(b *testing.B) {
	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendByte('x')
	})
}

func BenchmarkBuffer_AppendBytes(b *testing.B) {
	x := []byte("Hello, world! There is a normal bytes.")

	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendBytes(x)
	})
}

func BenchmarkBuffer_AppendFloat(b *testing.B) {
	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendFloat(1234.056789)
	})
}

func BenchmarkBuffer_AppendInt(b *testing.B) {
	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendInt(1234567890)
	})
}

func BenchmarkBuffer_AppendString(b *testing.B) {
	x := "Hello, world! There is a normal bytes."

	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendString(x)
	})
}

func BenchmarkBuffer_AppendTime(b *testing.B) {
	t := time.Now()

	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendTime(t, time.RFC1123)
	})
}

func BenchmarkBuffer_AppendUint(b *testing.B) {
	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendUint(1234567890)
	})
}

func BenchmarkBuffer_AppendRune(b *testing.B) {
	benchmarkBuffer(b, func(buf *Buffer) {
		buf.AppendRune('😀')
	})
}
