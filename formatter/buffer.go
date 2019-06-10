package formatter

import (
	"io"
	"fmt"
	"bytes"
)

// Simply wraps a bytes buffer to add extra writing capabilities
type Buffer struct {
	buffer *bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{
		buffer: bytes.NewBuffer([]byte("")),
	}
}

// Grabs the bytes Buffer directly
func (b *Buffer) GetBytesBuffer() *bytes.Buffer {
	return b.buffer
}

func (b *Buffer) WriteBytes(message []byte) {
	b.buffer.Write(message)
}

func (b *Buffer) Write(message string) {
	b.buffer.Write([]byte(message))
}

func (b *Buffer) Writef(message string, args ...interface{}) {
	b.Write(fmt.Sprintf(message, args...))
}

func (b *Buffer) Writeln(message string) {
	b.Write(message + "\n")
}

func (b *Buffer) WriteTo(out io.Writer) {
	b.buffer.WriteTo(out)
}

func (b *Buffer) String() string {
	return b.buffer.String()
}
