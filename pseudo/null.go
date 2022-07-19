package pseudo

import (
	. "github.com/snnyyp/ddWin/define"
	"io"
)

type Null struct {
	IfcePseudoBase
}

func (*Null) New() {}

func (*Null) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (*Null) Write(p []byte) (int, error) {
	return len(p), nil
}

func (*Null) Seek(int64, int) (int64, error) {
	return 0, nil
}

func (*Null) Close() error {
	return nil
}
