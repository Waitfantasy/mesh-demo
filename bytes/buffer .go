package bytes

import (
	"encoding/binary"
	"errors"
	"io"
)

var (
	NotEnoughErr = errors.New("没有足够的byte可以读取")
)

type Buffer struct {
	writeLen int
	readLen  int
	bytes    []byte
	tmp      []byte
	binOrder binary.ByteOrder
}

func NewBytesBuffer(size int, order binary.ByteOrder) *Buffer {
	return &Buffer{
		bytes:    make([]byte, size),
		tmp:      make([]byte, 8),
		binOrder: order,
	}
}

func (buf *Buffer) grow(n int) {
	// byte数组扩容
	byteLen := len(buf.bytes)
	if byteLen < buf.writeLen+n {
		newBytes := make([]byte, 2*byteLen+n)
		copy(newBytes, buf.bytes[:buf.writeLen])
		buf.bytes = newBytes
	}
}

func (buf *Buffer) WriteByte(v byte) {
	buf.grow(1)
	buf.bytes[buf.writeLen] = v
	buf.writeLen++
}

func (buf *Buffer) ReadByte() (v byte, err error) {
	if buf.readLen >= len(buf.bytes) {
		return 0, io.EOF
	}

	v = buf.bytes[buf.readLen]
	buf.readLen++
	return
}

func (buf *Buffer) WriteUint16(v uint16) {
	buf.grow(2)
	// 将v写入到bytes
	buf.binOrder.PutUint16(buf.bytes[buf.writeLen:], v)
	buf.writeLen += 2
}

func (buf *Buffer) ReadUint16() (v uint16, err error) {
	if !buf.enough(2) {
		return v, NotEnoughErr
	}

	v = buf.binOrder.Uint16(buf.bytes[buf.readLen : buf.readLen+2])
	buf.readLen += 2
	return v, nil
}

func (buf *Buffer) WriteUint64(v uint64) {
	buf.grow(8)
	buf.binOrder.PutUint64(buf.bytes[buf.writeLen:], v)
	buf.writeLen += 8
}

func (buf *Buffer) ReadUint64() (v uint64, err error) {
	if !buf.enough(8) {
		return v, NotEnoughErr
	}

	v = buf.binOrder.Uint64(buf.bytes[buf.readLen : buf.readLen+8])
	buf.readLen += 8
	return v, nil
}

func (buf *Buffer) enough(n int) bool {
	return buf.writeLen-buf.readLen >= n
}

func (buf *Buffer) Bytes() []byte {
	return buf.bytes[:buf.writeLen]
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.readLen >= len(b.bytes) {
		return 0, io.EOF
	}

	n = copy(p, b.bytes[b.readLen:])
	b.readLen += n
	return n, nil
}