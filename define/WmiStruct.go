package Define

type W32DiskDrive struct {
	DeviceID       string
	Model          string
	MediaType      string
	BytesPerSector uint32
	Size           uint64
	Partitions     uint32
	Manufacturer   string
	Index          uint32
}

type W32DiskPartition struct {
	DeviceID  string
	Caption   string
	DiskIndex uint32
	Name      string
}

type W32Partition struct {
	DeviceID string
}
