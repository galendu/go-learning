package main

import (
	"fmt"
	"time"
)

var (
	n int
)

func keyDay() {

	// 周六上课
	// 4 2
	// 5 1
	// 6 7
	// 0 6
	// 1 5
	// 2 4
	// 3 3
	// 今天周六还有7天
	now := time.Now()

	sub := 6 - int(now.Weekday())
	interval := sub

	if sub == 0 {
		interval = 7
	}

	firstStaurday := now.Add(24 * time.Duration(interval) * time.Hour)
	fmt.Println(firstStaurday.Format("2006-01-02"))

	for i := 0; i < 3; i++ {
		firstStaurday = firstStaurday.Add(24 * 7 * time.Hour)
		fmt.Println(firstStaurday.Format("2006-01-02"))
	}

}
