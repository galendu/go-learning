package main

import (
	"fmt"
	"time"
)

func timeFormat() {
	s1 := "1998-10-01 08:10:00"
	format1 := "2006-01-02 15:04:05"
	format2 := "20060102150405"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(format1, s1, loc)
	s2 := t.Format(format2)
	fmt.Println(s2)
}
