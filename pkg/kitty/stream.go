package kitty

import (
	"encoding/base64"
	"io"
)

const (
	chunkEncSize = 4096 //
	chunkRawSize = (chunkEncSize / 4) * 3
)

// streamPayload converts an image into a base64 encoded stream
type streamPayload struct {
	encoded [chunkEncSize]byte // buffer for encoded output
	raw     [chunkRawSize]byte // buffer for raw image input
	n       int                // number of bytes in raw buffer
	w       io.Writer          // output writer
}

func (spw *streamPayload) reset(w io.Writer) {
	spw.n = 0
	spw.w = w
}

// encode writes the next n bytes of encoded data to the output
func (spw *streamPayload) encode() error {
	base64.StdEncoding.Encode(spw.encoded[:], spw.raw[:spw.n])
	_, err := spw.w.Write(spw.encoded[:(spw.n+2)/3*4])
	spw.n = 0
	return err
}

func (spw *streamPayload) Write(b []byte) (n int, err error) {
	for len(b) > 0 {
		if spw.n == cap(spw.raw) {
			_, err = spw.w.Write([]byte("m=1;"))
			if err != nil {
				return
			}
			err = spw.encode()
			if err != nil {
				return
			}
			_, err = spw.w.Write([]byte("\033\\\033_G"))
			if err != nil {
				return
			}
		}

		l := copy(spw.raw[spw.n:], b)
		spw.n += l
		n += l
		b = b[l:]
	}

	return
}

func (spw *streamPayload) close() (err error) {
	if spw.n == 0 {
		_, err = spw.w.Write([]byte("m=0;\033\\"))
		return
	}
	_, err = spw.w.Write([]byte("m=0;"))
	if err != nil {
		return
	}
	err = spw.encode()
	if err != nil {
		return
	}
	_, err = spw.w.Write([]byte("\033\\"))
	return
}
