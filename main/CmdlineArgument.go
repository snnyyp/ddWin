package main

type CmdlineArgumentStruct struct {
	// 和dd对应的参数
	InputFile  string `arg:"--if" default:"*stdin*" help:"Read from a specific file instead of *stdin*"`  // 输入文件
	OutputFile string `arg:"--of" default:"*stdout*" help:"Write to a specific file instead of *stdout*"` // 输出文件
	BlockSize  int    `arg:"--bs" default:"512" help:"Size of a block"`                                   // 块大小，默认512B
	Count      int    `arg:"--count" default:"-1" help:"The amount of blocks to copy"`                    // 复制n块，默认为-1，即一直复制到文件末尾
	Skip       int    `arg:"--skip" default:"0" help:""`                                                  // 跳过输入文件的开头n块
	Seek       int    `arg:"--seek" default:"0" help:""`                                                  // 跳过输出文件的开头n块
	// ddWin的特殊参数
	// 若输出文件是分区，是否先卸载该分区
	Dismount bool `default:"false" help:"Whether dismount the partition before raw-writing to a partition"`
	// 是否在写入前先清空输出文件
	Erase bool   `default:"false" help:"Whether erase the output file initially"`
	List  string `help:""` // 列出所有的设备，包括分区、硬盘、伪设备
}

func (CmdlineArgumentStruct) Version() string {
	/*
		版本信息
	*/
	s := `
ddWin Ver0.1
Licensed under Apache 2.0
---
Author: snnyyp
Project Github Page: https://github.com/snnyyp/ddWin
E-mail: snnyyp_dev@outlook.com
`
	return s
}
