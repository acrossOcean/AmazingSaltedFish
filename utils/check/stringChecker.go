package check

import "AmazingSaltedFish/utils"

// StrChecker : 字符串检查器, 提供字符串相关检查
type StrChecker struct {
	value string

	options []StrCheckOption
}

// StrCheckOption : 字符串检查选项
type StrCheckOption struct {
	operator     StrOperator
	intTarget    int
	stringTarget string
	inStrTarget  []string
}

// StrOperator 字符串操作符
type StrOperator int

const (
	_ StrOperator = iota
	// StrOperatorLenLT : len(str) < ?
	StrOperatorLenLT
	// StrOperatorLenLE : len(str) <= ?
	StrOperatorLenLE
	// StrOperatorLenEQ : len(str) = ?
	StrOperatorLenEQ
	// StrOperatorLenNE : len(str) != ?
	StrOperatorLenNE
	// StrOperatorLenGT : len(str) > ?
	StrOperatorLenGT
	// StrOperatorLenGE : len(str) >= ?
	StrOperatorLenGE

	// StrOperatorStrEQ : str == ?
	StrOperatorStrEQ
	// StrOperatorStrNE : str != ?
	StrOperatorStrNE
	// StrOperatorStrIN : str in (?,?, ....)
	StrOperatorStrIN
)

// NewStrChecker 返回一个字符串检查器
func NewStrChecker(value string, options ...StrCheckOption) *StrChecker {
	result := &StrChecker{
		value: value,
	}
	result.options = make([]StrCheckOption, len(options))

	for i, o := range options {
		result.options[i] = o
	}

	return result
}

// NewStrCheckOption 返回一个字符串检查选项
func NewStrCheckOption(op StrOperator, intTarget int, strTarget string) StrCheckOption {
	var result = StrCheckOption{
		operator:     op,
		intTarget:    intTarget,
		stringTarget: strTarget,
	}

	return result
}

// NewStrCheckOptionInt 返回一个需要用到 int 对比的 检查选项
func NewStrCheckOptionInt(op StrOperator, intTarget int) StrCheckOption {
	var result = StrCheckOption{
		operator:  op,
		intTarget: intTarget,
	}

	return result
}

// NewStrCheckOptionStr 返回一个需要用到 string 对比的 检查选项
func NewStrCheckOptionStr(op StrOperator, strTarget string) StrCheckOption {
	var result = StrCheckOption{
		operator:     op,
		stringTarget: strTarget,
	}

	return result
}

// NewStrCheckOptionIN 返回一个 in 检查 的选项
func NewStrCheckOptionIN(strTargets []string) StrCheckOption {
	var result = StrCheckOption{
		operator:    StrOperatorStrIN,
		inStrTarget: strTargets,
	}

	return result
}

func (receiver *StrChecker) check() bool {
	hasFalse := false
	for _, o := range receiver.options {
		if hasFalse {
			return !hasFalse
		}

		switch o.operator {
		case StrOperatorLenLT:
			hasFalse = len(receiver.value) < o.intTarget == false
		case StrOperatorLenLE:
			hasFalse = len(receiver.value) <= o.intTarget == false
		case StrOperatorLenNE:
			hasFalse = len(receiver.value) != o.intTarget == false
		case StrOperatorLenEQ:
			hasFalse = len(receiver.value) == o.intTarget == false
		case StrOperatorLenGT:
			hasFalse = len(receiver.value) > o.intTarget == false
		case StrOperatorLenGE:
			hasFalse = len(receiver.value) >= o.intTarget == false
		case StrOperatorStrNE:
			hasFalse = receiver.value != o.stringTarget == false
		case StrOperatorStrEQ:
			hasFalse = receiver.value == o.stringTarget == false
		case StrOperatorStrIN:
			if len(o.inStrTarget) != len(utils.StringList(append(o.inStrTarget, receiver.value)).RemoveDuplicate()) {
				hasFalse = true
			}
		default:
			hasFalse = true
		}
	}

	return !hasFalse
}
