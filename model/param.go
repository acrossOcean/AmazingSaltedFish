package model

// ParamType 参数类型 分为 固定值和变量
type ParamType int

const (
	_ ParamType = iota
	// ParamTypeConst 固定值参数
	ParamTypeConst
	// ParamTypeVariable 变量参数(一般来自于其它 node/link 返回值)
	ParamTypeVariable
)

// ToInt 转换为 int 供检查器使用
func (receiver ParamType) ToInt() int {
	return int(receiver)
}
