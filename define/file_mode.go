package define

import "os"

const (
	// OBinary 二进制模式，Windows下必须
	OBinary = 32768
	// FileFlag rb+
	FileFlag = os.O_RDWR | OBinary | os.O_EXCL
	// FilePerm 缺省权限
	FilePerm = os.FileMode(511)
)
