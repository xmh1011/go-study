package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/shirou/gopsutil/mem"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				memInfo, err = mem.VirtualMemory()
				if err != nil {
					log.Fatal(err)
				}
				
				memoryUsage := 100 - memInfo.AvailablePercent
				
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				logLine := fmt.Sprintf("%s Memory Usage: %.2f%%\n", currentTime, memoryUsage)
				
				_, err = file.WriteString(logLine)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()
	
	// 监听中断信号，按下Ctrl+C时退出程序
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	
	ticker.Stop()
	done <- true
	fmt.Println("Program exited gracefully.")
}
