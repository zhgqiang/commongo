package data

// Base 未知类型
type Base struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// Bases 未知类型数组
type Bases []Base

// BaseInt 整型
type BaseInt struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// BaseInts 整型数组
type BaseInts []BaseInt

func (p BaseInts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BaseInts) Len() int {
	return len(p)
}

func (p BaseInts) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

// BaseInt8 8位整型
type BaseInt8 struct {
	Name  string `json:"name"`
	Value int8   `json:"value"`
}

// BaseInt8s 8位整型数组
type BaseInt8s []BaseInt8

func (p BaseInt8s) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BaseInt8s) Len() int {
	return len(p)
}

func (p BaseInt8s) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

// BaseInt32 32位整型
type BaseInt32 struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

// BaseInt32s 32位整型数组
type BaseInt32s []BaseInt32

func (p BaseInt32s) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BaseInt32s) Len() int {
	return len(p)
}

func (p BaseInt32s) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

// BaseInt64 64位整型
type BaseInt64 struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

// BaseInt64s 64位整型数组
type BaseInt64s []BaseInt64

func (p BaseInt64s) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BaseInt64s) Len() int {
	return len(p)
}

func (p BaseInt64s) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

// BaseFloat32 单精度浮点数
type BaseFloat32 struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

// BaseFloat32s 单精度浮点数数组
type BaseFloat32s []BaseFloat32

func (p BaseFloat32s) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BaseFloat32s) Len() int {
	return len(p)
}

func (p BaseFloat32s) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

// BaseFloat64 双精度浮点数
type BaseFloat64 struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// BaseFloat64s 双精度浮点数数组
type BaseFloat64s []BaseFloat64

func (p BaseFloat64s) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BaseFloat64s) Len() int {
	return len(p)
}

func (p BaseFloat64s) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}
