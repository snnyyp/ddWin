// Package DeviceCat /*
package DeviceCat

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/yusufpapurcu/wmi"

	Define "github.com/snnyyp/ddWin/Define"
)

func getBootloader(harddiskVolumePath string) string {
	disk, _ := os.Open(harddiskVolumePath)
	defer disk.Close()
	rawData := make([]byte, 512)
	_, _ = disk.Read(rawData)
	if rawData[450] == 238 {
		return "GPT"
	} else {
		return "MBR"
	}
}

func getInfos() (infos []Define.W32DiskDrive) {
	/*

	 */
	_ = wmi.Query("SELECT * FROM Win32_DiskDrive", &infos)
	return
}

func GetFormatDiskPartitionDeviceId(raw string) (formatStyle string) {
	/*

	 */
	pattern := regexp.MustCompile(`\d+`)
	matched := pattern.FindAllString(raw, -1)
	formatStyle = fmt.Sprintf("Disk #%v, Partition #%v", matched[0], matched[1])
	return
}

func GetPartitionContaining(harddiskVolumePath string) (partitionContaining []string) {
	/*

	 */
	var s []Define.W32DiskPartition
	var s2 []Define.W32Partition
	cmd := fmt.Sprintf("ASSOCIATORS OF {Win32_DiskDrive='%v'} WHERE AssocClass=Win32_DiskDriveToDiskPartition", harddiskVolumePath)
	_ = wmi.Query(cmd, &s)
	for _, i := range s {
		cmd := fmt.Sprintf("ASSOCIATORS OF {Win32_DiskPartition.DeviceID='%v'} WHERE AssocClass=Win32_LogicalDiskToPartition", GetFormatDiskPartitionDeviceId(i.Name))
		_ = wmi.Query(cmd, &s2)
		//必须在外循环内append slice，因为每次外循环s2都会改变
		for _, j := range s2 {
			partitionContaining = append(partitionContaining, j.DeviceID)
		}
	}
	return
}

func PrintDiskInfo() {
	/*

	 */
	fmt.Println("Available disks:")
	for _, info := range getInfos() {
		fmt.Printf("  %v:\n", info.DeviceID)
		fmt.Printf("    Model: %v\n", info.Model)
		fmt.Printf("    Bytes per sector: %v\n", info.BytesPerSector)
		fmt.Printf("    Media type: %v\n", info.MediaType)
		fmt.Printf("    Capability: %v\n", info.Size)
		fmt.Printf("    Bootloader: %v\n", getBootloader(info.DeviceID))
		fmt.Printf("    Partition amount: %v\n", info.Partitions)
		fmt.Printf("    Partition containing: %v\n", strings.Join(GetPartitionContaining(info.DeviceID), ", "))
		fmt.Println()
	}
}
