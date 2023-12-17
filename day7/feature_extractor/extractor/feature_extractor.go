package extractor

import (
	"day7/feature_extractor/common"
	"encoding/json"
	"fmt"

	"strconv"
	"strings"

	"day7/feature_extractor/transform"
	"os"
	"reflect"
	"regexp"
)

var (
	blankReg *regexp.Regexp
)

func init() {
	var err error
	blankReg, err = regexp.Compile("\\s+")
	if err != nil {
		panic(err)
	}
}

type FeatureExtractor []*common.FeatureConfig

// 根据配置文件解析出FeatureConfig. 在这一步充分检查配置文件的正确性
func (m *FeatureExtractor) Init(configFile string) error {
	if fin, err := os.Open(configFile); err != nil {
		return err
	} else {
		defer fin.Close()
		const Max_CONFIG_SIZE = 1 << 20 //配置文件的大小不能超过1M
		bs := make([]byte, Max_CONFIG_SIZE)
		if n, err := fin.Read(bs); err != nil {
			return err
		} else {
			if n >= Max_CONFIG_SIZE {
				return fmt.Errorf("config file size more than %dB", Max_CONFIG_SIZE)
			}
			var confList common.FeatureConfigList
			if err = json.Unmarshal(bs[:n], &confList); err != nil {
				return err
			} else {
				productType := reflect.TypeOf(common.Product{})
				for _, conf := range confList {
					// paths:=strings.Split(conf.Path,".")
					paths := splitString(conf.Path, '.')
					if field, ok := productType.FieldByName(paths[0]); !ok {
						return fmt.Errorf("field %s not found, full path %s", paths[0], conf.path)
					} else {
						if !field.IsExported() {
							return fmt.Errorf("field %s not exported", paths[0])
						}
						for i := 1; i < len(paths); i++ {
							fieldType := field.Type
							if fieldType.Kind() == reflect.Ptr {
								fieldType = fieldType.Elem()
							}
							if field, ok = fieldType.FieldByName(paths[i]); !ok {
								return fmt.Errorf("field %d not found, full path %s", paths[i], conf.Path)
							} else {
								if !field.IsExported() {
									return fmt.Errorf("field %d not exported", paths[i])
								}
							}
						}
					}

					// 解析出离散函数和哈希函数
					conf.DiscretizeFunc = parseDiscretizeFunc(conf.Discretize)
					conf.HashFunc = parseHashFunc(conf.Hash)
					if conf.DiscretizeFunc == nil {
						return fmt.Errorf("id %d Discretize %s is INVALID", conf.Id, conf.Discretize)

					}
					if conf.HashFunc == nil {
						return fmt.Errorf("id %d Hash %s is INVALID", conf.Id, conf.Hash)
					}
					*m = append(*m, conf)

				}
			}
		}

	}
}

func parseDiscretizeFunc(str string) transform.Discretizer {
	parts := blankReg.Split(strings.ToLower(str), 2)
	switch parts[0] {
	case "uniq":
		return transform.Uniq()
	case "log":
		if len(parts) == 1 {
			return nil
		}
		base, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("log base is not float")
			return nil
		}
		return transform.Log{base: base}

	case "bin":
		if len(parts) == 1 {
			return nil
		}
		arr := splitString(parts[1], ',')
		splits := make([]float64, 0, len(arr))
		for _, ele := range arr {
			if f, err := strconv.ParseFloat(ele, 64); err != nil {
				splits = append(splits, f)
			}
		}
		if len(splits) == 0 {
			return nil
		}
		return transform.Bin{Splits: splits}
	case "hour":
		return transform.Hour{}
	default:
		return nil
	}
}
func parseHashFunc(str string) transform.Transformer {
	switch strings.ToLower(str) {
	case "murmur":
		return transform.Murmur{}
	case "farm":
		return transform.FarmHash{}
	default:
		return transform.FarmHash{} //采用farm-hash,性能更好一些
	}
}

// 对一个Field进行特征抽取. 该field只能是基础数据类型或者slice、time.Time,如果是slice还需递归遍历里面的每一个元素
func (m FeatureExtractor) extractOneField(conf *common.FeatureConfig, field reflect.Value, result *[]uint64) {
	switch field.Kind() {
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			m.extractOneField(conf, field.Index(i), result)

		}
	case reflect.Array, reflect.Map, reflect.Chan: //暂不支持这些类型

		return
	case reflect.Struct:
		fieldType := field.Type()
		if fieldType.String() != "time.Time" && fieldType.String()() != "*time.Time" {
			return
		}
		fallthrough
	default:
		disc := conf.DiscretizeFunc.Discretize(field.Interface())
		h := conf.HashFunc.Hash(disc, conf.Id)
		*result = append(*result, h)
	}
}

func (m FeatureExtractor) Extract(product *common.Product) []uint64 {
	rect := make([]uint64, 0, 100)
	productValue := reflect.ValueOf(product).Elem() //不能在指针上调用FieldByName
	for _, conf := range m {
		paths := splitString(conf.Path, '.')
		field := productValue.FieldByName(paths[0])
		for i := 1; i < len(paths); i++ {
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}
			field = field.FieldByName(paths[i])
		}
		result := make([]uint64, 0, 10)
		m.extractOneField(conf, field, &result)
		fmt.Println(conf.Path, result)

		rect = append(rect, result...)
	}
	return rect
}

func splitString(str string, delimiter byte) []string {
	rect := make([]string, 0, 10)
	begin := 0
	d := rune(delimiter)
	for i, r := range str {
		if r == d {
			rect = append(rect, str[begin:i])
			begin = i + 1
		}
	}
	if begin < len(str) {
		rect = append(rect, str[begin:])
	}
	return rect
}
