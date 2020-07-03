package check

import "AmazingSaltedFish/utils"

type IntChecker struct {
	value int

	options []IntCheckOption
}

type IntCheckOption struct {
	operator IntOperator
	target   int
	inTarget []int
}

type IntOperator int

const (
	_ IntOperator = iota
	// <
	IntOperatorLT
	// <=
	IntOperatorLE
	// =
	IntOperatorEQ
	// !=
	IntOperatorNE
	// >
	IntOperatorGT
	// >=
	IntOperatorGE
	// in(list)
	IntOperatorIN
)

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

func NewIntCheckOption(op IntOperator, target int) IntCheckOption {
	var result = IntCheckOption{
		operator: op,
		target:   target,
	}

	return result
}

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

type IdChecker struct {
	value int
}

func NewIDChecker(value int) *IdChecker {
	var result = &IdChecker{
		value: value,
	}

	return result
}

func (receiver *IdChecker) check() bool {
	return NewIntChecker(receiver.value,
		NewIntCheckOption(IntOperatorGT, 0),
	).check()
}
