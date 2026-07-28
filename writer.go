package protorw

import (
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/proto"
)

// [WriteMsg] writes a protobuf message to the given [io.Writer].
func WriteMsg(w io.Writer, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	copy(buf[4:], data)

	_, err = w.Write(buf)
	return err
}

// [Writer] is a wrapper around an [io.Writer] that can write protobuf messages.
type Writer struct {
	w io.Writer
}

// [NewWriter] returns a new [Writer].
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// [Write] writes a protobuf message to the underlying [io.Writer].
func (w *Writer) Write(msg proto.Message) error {
	return WriteMsg(w.w, msg)
}
