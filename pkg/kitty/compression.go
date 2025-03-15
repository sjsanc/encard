package kitty

import (
	"compress/zlib"
	"io"
)

// zlibPayload compresses the streamPayload using zlib
type zlibPayload struct {
	buffer [16384]byte
	n      int
	spw    streamPayload
	zw     *zlib.Writer
}

// reset resets the zlibPayload with a new output writer
// The flag `o=z` indicates that we're using zlib compression
func (zp *zlibPayload) Reset(w io.Writer) {
	_, _ = w.Write([]byte("o=z,"))
	zp.spw.reset(w)

	// Creates a new zlib writer using the streamPayload writer
	zp.zw = zlib.NewWriter(&zp.spw)
	zp.n = 0
}

func (zp *zlibPayload) Write(b []byte) (n int, err error) {
	for len(b) > 0 {
		if zp.n == cap(zp.buffer) {
			_, err = zp.zw.Write(zp.buffer[:])
			if err != nil {
				return
			}
			zp.n = 0
		}
		m := copy(zp.buffer[zp.n:], b)
		zp.n += m
		n += m
		b = b[m:]
	}
	return
}

func (zp *zlibPayload) Close() error {
	if zp.n > 0 {
		if _, err := zp.zw.Write(zp.buffer[:zp.n]); err != nil {
			return err
		}
	}
	if err := zp.zw.Close(); err != nil {
		return err
	}
	return zp.spw.close()
}
