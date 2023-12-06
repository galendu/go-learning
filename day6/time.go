package main

import (
	"fmt"
	"time"
)

var (
	loc *time.Location
)

const (
	TIME_FMT = "2006-01-02 15:04:05"
	DATA_FMT = "20060102"
)

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

func basic() {
	t0 := time.Now()
	fmt.Printf("秒%d, 毫秒%d, 微秒%d, 纳秒%d\n", t0.Unix(), t0.UnixMilli(), t0.UnixMicro(), t0.UnixNano())
	time.Sleep(1 * time.Second)

	t1 := time.Now()
	diff1 := t1.Sub(t0)     // 计算时间差,返回类型是time.Duration
	diff2 := time.Since(t0) // 等价于t1:=time.Now();t1.Sub(t0)
	fmt.Println(diff1, diff2)
	fmt.Printf("diff1=diff2 %t\n", int(diff1.Seconds()) == int(diff2.Seconds())) //在秒级别上diff1和diff2是相等的
	fmt.Printf("t1>t0 %t\n", t1.After(t0))                                       //判断时间的先后顺序

	d := time.Duration(3 * time.Hour) //Duration表示两个时刻之间的距离
	fmt.Println(d)
	t2 := t1.Add(d)
	fmt.Printf("hour1 %d, hour2 %d\n", t1.Hour(), t2.Hour()) //获取一个时刻的Hour
	fmt.Printf("week day %s %d\n", t2.Weekday().String(), t2.Weekday())
	fmt.Printf("year %d\n", t2.Year())
	fmt.Printf("month %s %d %d\n", t2.Month(), t2.Month(), int(t2.Month())) //在go语言内部Month就是int类型。type Month int
	fmt.Printf("day %d\n", t2.Day())
	fmt.Printf("day in year %d\n", t2.YearDay())

}

// 时间格式化
func parse_format() {
	now := time.Now()
	ts := now.Format(TIME_FMT)
	fmt.Println("格式化时间", ts)

	t, _ := time.Parse(TIME_FMT, ts)
	if t.Unix() != now.Unix() {
		fmt.Printf("Parse时间解析错误: %d!=%d\n", t.Unix(), now.Unix())
	}
	t, _ = time.ParseInLocation(TIME_FMT, ts, loc)
	if t.Unix() != now.Unix() {
		fmt.Printf("ParseInLocation时间解析错误: %s\n", t.Format(TIME_FMT))
	}

}

// 周期执行任务
func ticker() {
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		<-tk.C
		fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	}
	tk.Stop()
}

// 经过一段时间后,触发一次任务
func timer() {
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tm := time.NewTimer(3 * time.Second)
	<-tm.C
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tm.Stop()

	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	<-time.After(3 * time.Second)
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
}

func main() {
	// basic()
	// parse_format()
	timer()
}
