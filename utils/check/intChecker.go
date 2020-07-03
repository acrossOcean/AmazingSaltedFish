package check

import "AmazingSaltedFish/utils"

// IntChecker : 数字检查器, 提供数字相关检查
type IntChecker struct {
	value int

	options []IntCheckOption
}

// IntCheckOption : int 检查选项
type IntCheckOption struct {
	operator IntOperator
	target   int
	inTarget []int
}

// IntOperator int 操作符
type IntOperator int

const (
	_ IntOperator = iota
	// IntOperatorLT : <
	IntOperatorLT
	// IntOperatorLE : <=
	IntOperatorLE
	// IntOperatorEQ : =
	IntOperatorEQ
	// IntOperatorNE : !=
	IntOperatorNE
	// IntOperatorGT : >
	IntOperatorGT
	// IntOperatorGE : >=
	IntOperatorGE
	// IntOperatorIN : in(list)
	IntOperatorIN
)

// NewIntChecker 返回一个 int 检查器
func NewIntChecker(value int, options ...IntCheckOption) *IntChecker {
	result := &IntChecker{
		value: value,
	}
	result.options = make([]IntCheckOption, len(options))

	for i, o := range options {
		result.options[i] = o
	}

	return result
}

// NewIntCheckOption 返回一个 int 检查选项
func NewIntCheckOption(op IntOperator, target int) IntCheckOption {
	var result = IntCheckOption{
		operator: op,
		target:   target,
	}

	return result
}

// NewIntCheckOptionIN 返回一个 in 检查 选项
func NewIntCheckOptionIN(targets []int) IntCheckOption {
	var result = IntCheckOption{
		operator: IntOperatorIN,
		inTarget: targets,
	}

	return result
}

func (receiver *IntChecker) check() bool {
	hasFalse := false
	for _, o := range receiver.options {
		if hasFalse {
			return !hasFalse
		}
		switch o.operator {
		case IntOperatorLT:
			hasFalse = receiver.value < o.target == false
		case IntOperatorLE:
			hasFalse = receiver.value <= o.target == false
		case IntOperatorNE:
			hasFalse = receiver.value != o.target == false
		case IntOperatorEQ:
			hasFalse = receiver.value == o.target == false
		case IntOperatorGT:
			hasFalse = receiver.value > o.target == false
		case IntOperatorGE:
			hasFalse = receiver.value >= o.target == false
		case IntOperatorIN:
			if len(o.inTarget) != len(utils.IntList(append(o.inTarget, receiver.value)).RemoveDuplicate()) {
				hasFalse = true
			}
		default:
			hasFalse = true
		}
	}

	return true
}

// IDChecker ID检查器, 检查id的合法性
type IDChecker struct {
	value int
}

// NewIDChecker 返回一个 id 检查器
func NewIDChecker(value int) *IDChecker {
	var result = &IDChecker{
		value: value,
	}

	return result
}

func (receiver *IDChecker) check() bool {
	return NewIntChecker(receiver.value,
		NewIntCheckOption(IntOperatorGT, 0),
	).check()
}
