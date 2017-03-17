type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}

// ReadWriter stores pointers to a Reader and a Writer.
type ReadWriter struct {
    *Reader
    *Writer
}
