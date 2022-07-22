package DeviceCat

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/0xrawsec/golang-win32/win32/kernel32"
	psutil "github.com/shirou/gopsutil/disk"

	Define "github.com/snnyyp/ddWin/Define"
)

func GetHarddiskVolume() (out []string) {
	/*

	 */
	out = make([]string, 0)
	devSlice, _ := kernel32.QueryDosDevice("")
	for _, dev := range devSlice {
		if IsHarddiskVolume(dev) {
			out = append(out, dev)
		}
	}
	return
}

func GetAliasPaths(harddiskVolume string) (out []string) {
	/*

	 */
	out = make([]string, 0)
	devList, _ := kernel32.QueryDosDevice("")
	for _, dev := range devList {
		r, _ := kernel32.QueryDosDevice(dev)
		fmtPath := fmt.Sprintf(`\Device\%v`, harddiskVolume)
		if len(r) > 0 && r[0] == fmtPath && dev != harddiskVolume {
			out = append(out, `\\.\`+dev)
		}
	}
	return
}

func GetDriveLetter(harddiskVolume string) (out string) {
	for _, alias := range GetAliasPaths(harddiskVolume) {
		if IsDriveLetterPath(alias) {
			out = alias[4:]
		}
	}
	return
}

func IsPartitionRecognizable(harddiskVolume string) bool {
	return GetDriveLetter(harddiskVolume) != ""
}

func GetFileSystem(driveLetter string) (fileSystem string) {
	stats, _ := psutil.Partitions(false)
	for _, stat := range stats {
		if stat.Device == driveLetter {
			fileSystem = stat.Fstype
		}
	}
	return
}

func GetUsage(driveLetter string) (total uint64, free uint64, used uint64, usedPercent float64) {
	usage, _ := psutil.Usage(driveLetter)
	total = usage.Total
	free = usage.Free
	used = usage.Used
	usedPercent = usage.UsedPercent
	return
}

func GetCapability(harddiskVolumePath string) int64 {
	/*

	 */
	file, _ := os.Open(harddiskVolumePath)
	defer file.Close()
	fd := syscall.Handle(file.Fd())
	s := Define.GetLengthInfo{}
	bytesReturned := uint32(0)
	_ = syscall.DeviceIoControl(fd, Define.IoctlDiskGetLengthInfo,
		nil, 0,
		(*byte)(unsafe.Pointer(&s)), uint32(unsafe.Sizeof(s)), &bytesReturned,
		nil)
	return s.Length
}

func GetDetectedFileSystem(harddiskVolumePath string) (fileSystem string) {
	fileSystem = "Unknown" //默认是未知的文件系统
	harddiskVolume, _ := os.Open(harddiskVolumePath)
	defer harddiskVolume.Close()
	//判断Ext
	harddiskVolume.Seek(0, 0)
	data := make([]byte, 2048)
	if _, _ = harddiskVolume.Read(data); data[1080] == 83 && data[1081] == 239 {
		fileSystem = "Ext"
	}
	return
}

func PrintPartitionInfo() {
	/*

	 */
	fmt.Println("Available partitions:")
	for _, harddiskVolume := range GetHarddiskVolume() {
		fmt.Printf("  \\\\.\\%s:\n", harddiskVolume)
		fmt.Printf("    Alias paths: %s\n", strings.Join(GetAliasPaths(harddiskVolume), ", "))
		if IsPartitionRecognizable(harddiskVolume) {
			fmt.Println("    System recognizable: True")
			driveLetter := GetDriveLetter(harddiskVolume)
			fmt.Printf("    Drive letter: %v\n", driveLetter)
			fmt.Printf("    File system: %v\n", GetFileSystem(driveLetter))
			total, free, used, usedPercent := GetUsage(driveLetter)
			fmt.Printf("    Total: %v\n", total)
			fmt.Printf("    Free: %v\n", free)
			fmt.Printf("    Used: %v\n", used)
			fmt.Printf("    Used percent: %v\n", usedPercent)
		} else {
			fmt.Println("    System recognizable: False")
			harddiskVolumePath := `\\.\` + harddiskVolume
			fmt.Printf("    Detected filesystem: %v\n", GetDetectedFileSystem(harddiskVolumePath))
			fmt.Printf("    Capability: %v\n", GetCapability(harddiskVolumePath))
		}
		fmt.Println()
	}
}
