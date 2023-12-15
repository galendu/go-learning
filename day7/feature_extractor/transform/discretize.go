package transform

type Discretizer interface {
	//把特征取值转成字符串
	Discretize(i interface{}) string
}
