package Define

import "os"

const (
	O_BINARY           = 32768
	InputFileOpenMode  = os.O_RDONLY | O_BINARY | os.O_EXCL
	OutputFileOpenMode = os.O_WRONLY | O_BINARY | os.O_EXCL
	UniversalOpenMode  = os.O_RDWR | O_BINARY | os.O_EXCL // rb+
	DefaultPermission  = os.FileMode(511)
)
