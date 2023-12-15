package common

import (
	"day7/feature_extractor/transform"
	"time"
)

// feature 特征 特色

type Location struct {
	Province string // 省
	City     string
}

type User struct {
	Name    string
	Age     int
	Gender  byte // 性别
	Address *Location
}

type Product struct {
	Id        int
	Name      string
	Sales     int       //销量
	GoodRate  float32   //好评率
	Business  *User     //商家
	ShelfTime time.Time //上架时间
	Tags      []string
}

type FeatureConfig struct {
	Id             int                   `json:"id"`
	Path           string                `json:"path"`
	Discretize     string                `json:"discretize"` //离散化
	Hash           string                `json:"hash"`
	DiscretizeFunc transform.Discretizer `json:"-"`
	HashFunc       transform.Transformer `json:"-"`
}

type FeatureConfigList []*FeatureConfig
