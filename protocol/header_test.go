package protocol

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/shein/bytes"
	"io"
	"testing"
)

func TestNewHeader(t *testing.T) {
	h := NewHeader(Request, true, 1, 4, 2349789)
	buf := bytes.NewBytesBuffer(HeaderSize, binary.BigEndian)
	buf.WriteUint16(h.MagicNum)
	buf.WriteByte(h.Features)
	buf.WriteByte(h.Version)
	buf.WriteByte(h.Status)
	buf.WriteByte(h.Serialize)
	buf.WriteByte(h.Reserve)
	buf.WriteUint64(h.RequestId)
	fmt.Println(buf.Bytes())
	temp := make([]byte, HeaderSize, HeaderSize)
	_, err := io.ReadAtLeast(bufio.NewReader(buf), temp, HeaderSize)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(h)
	fmt.Println(binary.BigEndian.Uint16(temp[:2]))
	fmt.Println(temp[2])
	fmt.Println(temp[3])
	fmt.Println(temp[4])
	fmt.Println(temp[5])
	fmt.Println(temp[6])
	fmt.Println(binary.BigEndian.Uint64(temp[7:]))
	fmt.Println(temp)
}
