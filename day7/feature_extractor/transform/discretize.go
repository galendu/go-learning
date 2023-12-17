package transform

import (
	"math"
	"strconv"
	"time"
)

type Discretizer interface {
	//把特征取值转成字符串
	Discretize(i interface{}) string
}

type Uniq struct{}

func (Uniq) Discretize(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.Itoa(int(v))
	case float64:
		return strconv.Itoa(int(v))
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

type Log struct {
	Base float64
}

func (self Log) Discretize(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case int32:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case int64:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case float32:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case float64:
		return strconv.Itoa(int(self._log(v, self.Base)))
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""

	}
}

func (Log) _log(f float64, base float64) float64 {
	if f < 1 {
		f = 1
	}
	return math.Log(f) / math.Log(base)
}

type Bin struct {
	Splits []float64
}

func (self Bin) Discretize(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))

	case int32:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case int64:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case float32:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

func (Bin) BinDiscretize(f float64, splits []float64) int {
	index := 0
	for _, split := range splits {
		if f < split {
			break
		}
		index++

	}
	return index
}

type Hour struct{}

func (Hour) Discretize(i interface{}) string {
	switch v := i.(type) {
	case time.Time:
		return strconv.FormatInt(int64(v.Hour()), 10)
	case int64:
		return strconv.FormatInt(int64(time.Unix(v, 0).Hour()), 10)
	default:
		return ""
	}
}
