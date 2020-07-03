package model

type ParamType int

const (
	_ ParamType = iota
	// 固定值参数
	ParamTypeConst
	// 变量参数(一般来自于其它 node/link 返回值)
	ParamTypeVariable
)

func (receiver ParamType) ToInt() int {
	return int(receiver)
}
