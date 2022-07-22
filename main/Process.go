package main

import "io"

func Process() (string, error) {
	// 最后关闭文件
	defer inputFileHandle.Close()
	defer outputFileHandle.Close()
	// 先移动好输入、输出文件的指针
	_, err := inputFileHandle.Seek(int64(argStruct.Skip*argStruct.BlockSize), io.SeekStart)
	if err != nil {
		return "An unexpected error occurred while skipping input file", err
	}
	_, err = outputFileHandle.Seek(int64(argStruct.Seek*argStruct.BlockSize), io.SeekStart)
	if err != nil {
		return "An unexpected error occurred while seeking output file", err
	}
	// 声明缓冲区
	buf := make([]byte, argStruct.BlockSize)
	// 声明计数器
	var n, c int
	// 写入
	// argStruct.Count < 0，说明是无限读入
	for c < argStruct.Count || argStruct.Count < 0 {
		n, err = inputFileHandle.Read(buf)
		totalInput += n // 总输入增加
		if err != nil { // 读取发生异常
			// 不用担心EOF异常，会做处理的
			return "An unexpected error occurred while reading from input file", err
		}
		// 读取未发生异常，开始写入
		n, err = outputFileHandle.Write(buf[:n])
		totalOutput += n // 总输出增加
		if err != nil {  // 写入发生异常
			return "An unexpected error occurred while writing to output file", err
		}
		// 写入未发生异常
		if argStruct.Count > 0 {
			c++
		}
	}
	return "", nil
}
