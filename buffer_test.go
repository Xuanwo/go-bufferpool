package bufferpool

import (
	"bytes"
	"testing"
	"time"
)

func testBuffer(t *testing.T, fn func(buf *Buffer)) {
	p := New(1024)

	t.Run("", func(t *testing.T) {
		buf := p.Get()
		fn(buf)
		buf.Free()
	})
}

func TestBuffer_AppendBool(t *testing.T) {
	testBuffer(t, func(buf *Buffer) {
		buf.AppendBool(true)

		if !bytes.Equal(buf.Bytes(), []byte("true")) {
			t.Errorf("expect %s, actual %s", "true", buf.Bytes())
		}
	})
}

func TestBuffer_AppendByte(t *testing.T) {
	testBuffer(t, func(buf *Buffer) {
		buf.AppendByte('x')

		if !bytes.Equal(buf.Bytes(), []byte("x")) {
			t.Errorf("expect %s, actual %s", "x", buf.Bytes())
		}
	})
}

func TestBuffer_AppendBytes(t *testing.T) {
	testBuffer(t, func(buf *Buffer) {
		x := "Hello, world! There is a normal bytes."

		buf.AppendBytes([]byte(x))

		if !bytes.Equal(buf.Bytes(), []byte(x)) {
			t.Errorf("expect %s, actual %s", "x", buf.Bytes())
		}
	})
}

func TestBuffer_AppendFloat(t *testing.T) {
	testBuffer(t, func(buf *Buffer) {
		buf.AppendFloat(1234.056789)
	})
}

func TestBuffer_AppendInt(t *testing.T) {
	testBuffer(t, func(buf *Buffer) {
		buf.AppendInt(1234567890)
	})
}

func TestBuffer_AppendString(t *testing.T) {
	x := "Hello, world! There is a normal bytes."

	testBuffer(t, func(buf *Buffer) {
		buf.AppendString(x)
	})
}

func TestBuffer_AppendUint(t *testing.T) {
	ti := time.Now()

	testBuffer(t, func(buf *Buffer) {
		buf.AppendTime(ti, time.RFC1123)
	})
}

func TestBuffer_AppendRune(t *testing.T) {
	testBuffer(t, func(buf *Buffer) {
		buf.AppendRune('😀')
	})
}
