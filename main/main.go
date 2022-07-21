package main

import (
	"fmt"
	"github.com/snnyyp/ddWin/list"
	"github.com/yusufpapurcu/wmi"
)

func getDiskDrivers() {
	type Win32_DiskDrive struct {
		Caption      string
		Name         string
		DeviceID     string
		Model        string
		Index        int
		Partitions   int
		Size         int
		PNPDeviceID  string
		Status       string
		SerialNumber string
		Manufacturer string
		MediaType    string
		Description  string
		SystemName   string
	}

	var dst []Win32_DiskDrive

	//query := wmi.CreateQuery(&dst, "")
	//if err := wmi.Query(query, &dst); err != nil {
	//	fmt.Println(err.Error())
	//}

	query := wmi.CreateQuery(&dst, "")
	wmi.Query(query, &dst)
	for _, v := range dst {
		fmt.Printf("%+v\n\n", v)
	}
}

func main() {
	//getDiskDrivers()
	list.ListPseudo()
}
