package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

const lowerHex = "0123456789abcdef"

type Writer struct {
	io.Writer
	c int
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	buf := []byte(`\x00`)
	var b byte

	for n, b = range p {
		buf[2] = lowerHex[b/16]
		buf[3] = lowerHex[b%16]
		w.Writer.Write(buf)
		w.c++
	}

	n++
	return
}

func writeCode(files []FileInfo) error {
	fd, err := os.Create("./bindata.go")
	if err != nil {
		return err
	}
	defer fd.Close()

	w := bufio.NewWriter(fd)
	defer w.Flush()

	if _, err = fmt.Fprintf(w, "package main\n\n"); err != nil {
		return err
	}

	if err := writeHeader(w); err != nil {
		return err
	}

	for i := range files {
		if err := writeBytes(w, &files[i]); err != nil {
			return err
		}
	}

	return nil
}

func writeBytes(w io.Writer, file *FileInfo) error {
	r, err := os.Open(file.Path)
	if err != nil {
		return err
	}
	defer r.Close()

	if _, err := fmt.Fprintf(w, `var %s = "`, file.Name); err != nil {
		return err
	}

	gz := gzip.NewWriter(&Writer{Writer: w})
	_, err = io.Copy(gz, r)
	gz.Close()

	return err
}

func writeHeader(w io.Writer) error {
	_, err := fmt.Fprintf(w, `import (
	"bytes"
	"compress/gzip"
	"io"
)

func asset(bs []byte) []byte {
	var b bytes.Buffer
	gz, _ := gzip.NewReader(bytes.NewReader(bs))
	io.Copy(&b, gz)
	gz.Close()
	return b.Bytes()
}

`)
	return err
}
