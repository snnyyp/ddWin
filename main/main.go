package main

import (
	"fmt"
	goarg "github.com/alexflint/go-arg"
	"github.com/snnyyp/ddWin/Define"
	"github.com/snnyyp/ddWin/DeviceCat"
	"github.com/snnyyp/ddWin/Pseudo"
	"io"
	"os"
)

/*
--if，输入文件，必须存在，默认为*stdin*
--of，输出文件，可以不存在，不存在则创建，默认为*stdout*
--bs，设置输入/输出块的大小
--count，仅拷贝n个块
--skip，跳过输入文件开头的n个块
--seek，跳过输出文件开头的n个块
--dismount 卸载分区，在写分区的时候要启用
--list 列出所有的硬盘、分区、伪设备的信息
--erase 在写入前清空输出文件
--version 打印版本信息
*/

var (
	argStruct        = CmdlineArgumentStruct{}
	inputFileHandle  Define.IfceBasicFileHandle
	outputFileHandle Define.IfceBasicFileHandle
)

func main() {
	// 解析命令行参数
	goarg.MustParse(&argStruct)
	fmt.Printf("%#v\n", argStruct)
	// 解析list参数
	switch argStruct.List {
	case "all":
		DeviceCat.PrintDiskInfo()
		DeviceCat.PrintPartitionInfo()
		DeviceCat.PrintPseudoInfo()
		return
	case "disk":
		DeviceCat.PrintDiskInfo()
		return
	case "partition":
		DeviceCat.PrintPartitionInfo()
		return
	case "pseudo":
		DeviceCat.PrintPseudoInfo()
		return
	case "": // 未传入
	default: // 传入了错误的参数
		fmt.Println("Unknown --list argument:", argStruct.List)
		return
	}
	// 设置Ctrl+C处理
	setupCtrlCHandler()
	// 清空输出文件；或输出文件不存在，则创建之
	if argStruct.Erase || !IsOutputFileExist() {
		EraseOrTouchOutputFile()
	}
	// 打开输入文件
	if Pseudo.IsPseudo(argStruct.InputFile) { // 如果是伪设备
		inputFileHandle = Pseudo.GetByName(argStruct.InputFile)
	} else { // 正常打开文件
		inputFileHandle, _ = os.OpenFile(argStruct.InputFile,
			Define.UniversalOpenMode, Define.DefaultPermission)
	}
	// 打开输出文件
	if Pseudo.IsPseudo(argStruct.OutputFile) { // 如果是伪设备
		outputFileHandle = Pseudo.GetByName(argStruct.OutputFile)
	} else { // 正常打开文件
		outputFileHandle, _ = os.OpenFile(argStruct.OutputFile,
			Define.UniversalOpenMode, Define.DefaultPermission)
	}
	// 卸载分区（输出文件）
	if argStruct.Dismount {
		DismountOutputPartition()
	}
	// 开始正式处理
	description, err := Process()
	if err != nil && err != io.EOF { // 发生了除EOF以外的异常
		fmt.Println(description)
		fmt.Println(err)
	}
	// 打印统计信息
	printAnalysis()
}
