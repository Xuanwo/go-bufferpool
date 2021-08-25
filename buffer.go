package bufferpool

import (
	"strconv"
	"time"
	"unicode/utf8"
)

type Buffer struct {
	bs   []byte
	pool *Pool
}

// AppendBool will append a bool into buffer.
//
// true => "true"
// false => "false"
func (b *Buffer) AppendBool(v bool) {
	b.bs = strconv.AppendBool(b.bs, v)
}

// AppendByte will append a byte into buffer.
func (b *Buffer) AppendByte(v byte) {
	b.bs = append(b.bs, v)
}

// AppendBytes will append an bytes slice into buffer.
func (b *Buffer) AppendBytes(v []byte) {
	b.bs = append(b.bs, v...)
}

// AppendFloat will append a float64 into buffer.
func (b *Buffer) AppendFloat(f float64) {
	b.bs = strconv.AppendFloat(b.bs, f, 'f', -1, 64)
}

// AppendInt will append an int64 into buffer.
func (b *Buffer) AppendInt(i int64) {
	b.bs = strconv.AppendInt(b.bs, i, 10)
}

// AppendString will append a string into buffer.
func (b *Buffer) AppendString(s string) {
	b.bs = append(b.bs, s...)
}

// AppendTime will append time with input layout.
func (b *Buffer) AppendTime(t time.Time, layout string) {
	b.bs = t.AppendFormat(b.bs, layout)
}

// AppendUint will append a uint64 into buffer.
func (b *Buffer) AppendUint(i uint64) {
	b.bs = strconv.AppendUint(b.bs, i, 10)
}

// AppendRune will append a rune into buffer.
func (b *Buffer) AppendRune(r rune) {
	rs := make([]byte, 4)
	size := utf8.EncodeRune(rs, r)
	b.bs = append(b.bs, rs[:size]...)
}

// Bytes will return underlying bytes.
func (b *Buffer) Bytes() []byte {
	return b.bs
}

// BytesCopy will return copied underlying bytes.
func (b *Buffer) BytesCopy() []byte {
	bs := make([]byte, len(b.bs))
	copy(bs, b.bs)
	return bs
}

// String will convert underlying bytes to string and return.
func (b *Buffer) String() string {
	return string(b.bs)
}

// Free will free underlying bytes and put into pool.
//
// After free, this buffer should not be touched anymore.
func (b *Buffer) Free() {
	b.Reset()
	b.pool.p.Put(b)
}

// Reset will reset underlying buffer.
func (b *Buffer) Reset() {
	b.bs = b.bs[:0]
}

// Write implements io.Writer to support write into buffer.
func (b *Buffer) Write(bs []byte) (int, error) {
	b.bs = append(b.bs, bs...)
	return len(bs), nil
}
