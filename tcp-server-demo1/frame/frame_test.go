package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"
)

func TestEncode(t *testing.T) {
	codec := NewMyFrameCodec()
	buf := make([]byte, 0, 128)
	rw := bytes.NewBuffer(buf)

	err := codec.Encode(rw, []byte("hello"))
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	var totalLen int32
	err = binary.Read(rw, binary.BigEndian, &totalLen)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	if totalLen != 9 {
		t.Errorf("wnat 9, actual %d", totalLen)
	}

	left := rw.Bytes()
	if string(left) != "hello" {
		t.Errorf("want hello, actual %s", string(left))
	}
}

func TestDecode(t *testing.T) {
	codec := NewMyFrameCodec()
	data := []byte{0x0, 0x0, 0x0, 0x9, 'h', 'e', 'l', 'l', 'o'}

	payload, err := codec.Decode(bytes.NewReader(data))
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	if string(payload) != "hello" {
		t.Errorf("want hello, actual %s", string(payload))
	}
}

// Add for test if case

type ReturnErrorWriter struct {
	W io.Writer
	Wn int
	wc int
}

func (w *ReturnErrorWriter) Write(p []byte) (n int, err error) {
	w.wc++
	if w.wc >= w.Wn {
		return 0, errors.New("write error")
	}
	return w.W.Write(p)
}

type ReturnErrorReader struct {
	R io.Reader
	Rn int
	rc int
}

func (r *ReturnErrorReader) Read(p []byte) (n int, err error) {
	r.rc++
	if r.rc >= r.Rn {
		return 0, errors.New("read error")
	}

	return r.R.Read(p)
}

func TestEncodeWithWriteFail(t *testing.T) {
	codec := NewMyFrameCodec()
	buf := make([]byte, 0, 128)
	rw := bytes.NewBuffer(buf)

	err := codec.Encode(&ReturnErrorWriter{W: rw, Wn: 1}, []byte("hello"))
	if err == nil {
		t.Errorf("want write error, actual nil")
	}

	err = codec.Encode(&ReturnErrorWriter{W: rw, Wn: 2}, []byte("hello"))
	if err == nil {
		t.Errorf("want write error, actual nil")
	}
}

func TestDeodeWithReadFail(t *testing.T) {
	codec := NewMyFrameCodec()
	data := []byte{0x0, 0x0, 0x0, 0x9, 'h', 'e', 'l', 'l', 'o'}

	_, err := codec.Decode(&ReturnErrorReader{R: bytes.NewReader(data), Rn: 1})
	if err == nil {
		t.Errorf("want read error, actual nil")
	}

	_, err = codec.Decode(&ReturnErrorReader{R: bytes.NewReader(data), Rn: 2})
	if err == nil {
		t.Errorf("want read error, actual nil")
	}
}