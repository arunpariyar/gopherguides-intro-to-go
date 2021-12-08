package cli

import (
	"io"
	"os"
)

type IO struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func (oi IO) Stdout() io.Writer {
	if oi.Out == nil {
		return os.Stdout
	}
	return oi.Out
}
