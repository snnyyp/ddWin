package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 统计数据
var (
	totalInput  = 0
	totalOutput = 0
	startTime   = time.Now()
)

func setupCtrlCHandler() {
	sigChan := make(chan os.Signal, 1)
	// 等待信号
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		// 关闭输入输出文件
		inputFileHandle.Close()
		outputFileHandle.Close()
		// 打印统计信息
		printAnalysis()
		// 直接退出，否则会程序仍会继续运行
		os.Exit(0)
	}()
}

func printAnalysis() {
	/*
		打印统计信息
	*/
	fmt.Println() // 先换个行
	fmt.Printf("Recorded %v+%v in\n",
		totalInput/argStruct.BlockSize, totalInput%argStruct.BlockSize)
	fmt.Printf("Recorded %v+%v out\n",
		totalOutput/argStruct.BlockSize, totalOutput%argStruct.BlockSize)
	fmt.Printf("%vBytes copied\n",
		totalInput+totalOutput)
	t := time.Since(startTime)
	fmt.Printf("Time consumed: %vSec (%vMin)\n",
		t.Seconds(), t.Minutes())
	fmt.Printf("Average copy speed: %vKB/s\n",
		float64((totalInput+totalOutput)/1024)/t.Seconds())
}
