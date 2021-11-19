package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t)

	t1 := time.Now().Unix()
	fmt.Println(t1)

	day := time.Now().Format("2006-01-02")
	fmt.Println(day)

	start := day + " 00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", start, time.Local)
	fmt.Println(startTime)
	fmt.Println(startTime.Unix())

	end := day + " 23:59:59"
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end, time.Local)
	fmt.Println(endTime)
	fmt.Println(endTime.Unix())
}
