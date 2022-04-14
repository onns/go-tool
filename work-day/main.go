package main

import (
	"fmt"
	"time"
)

/*
@Time : 2021/10/10 19:49
@Author : onns
@File : work-day/main.go
*/

func main() {
	now := time.Now()
	n := 0
	for {
		if now.Year() != time.Now().Year() {
			break
		}
		now = now.AddDate(0,0,1)
		if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
			continue
		}
		n += 1
	}
	fmt.Printf("今年还需上班 %d 天\n", n)
}