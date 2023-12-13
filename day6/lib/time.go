package main

import (
	"fmt"
	"time"
)

func time_resolve() {
	TIME_FMT := "2006-01-02 15:04:05"
	now := time.Now()
	ts := now.Format(TIME_FMT)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(TIME_FMT, ts, loc)
	fmt.Println(ts, t)
}

func timeCalculations() {
	t0 := time.Now()
	fmt.Printf("unix: %v, unixMilli: %v, unixMicro: %v,UnixNano: %v\n", t0.Unix(), t0.UnixMilli(), t0.UnixMicro(), t0.UnixNano())
	t2 := time.Now().Add(time.Hour * 12)
	fmt.Println(t2.Year(), t2.Month(), t2.Day(), t2.YearDay(), t2.Weekday().String(), t2.Weekday(), t0.Hour(), t0.Minute(), t0.Second())

	time.Sleep(2 * time.Second)
	t1 := time.Now().Add(time.Hour * 2)

	diff1 := t1.Sub(t0)                   //计算t1跟t0的时间差,返回类型是time.Duration
	diff2 := time.Since(t0)               //计算当前时间跟t0的时间差,返回类型是time.Duration
	diff3 := time.Duration(3 * time.Hour) //Duration表示两个时刻之间的距离
	t4 := t0.Add(diff3)

	fmt.Printf("diff1: %s, diff2: %s, diff3: %s, t4: %s,t4在t0之后: %t\n", diff1, diff2, diff3, t4, t4.After(t0))
}
func scheduledExecution() {
	tm := time.NewTimer(3 * time.Second)
	<-tm.C //阻塞3秒钟
	//do something
	tm.Stop()

	//或者用:
	<-time.After(3 * time.Second) //阻塞3秒钟

}

func periodicalExecution() {

	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 6; i++ {
		<-tk.C //阻塞一秒钟
		// do something
	}
	tk.Stop()
}
func main1() {
	// timeCalculations()
	// scheduledExecution()
	periodicalExecution()
}
