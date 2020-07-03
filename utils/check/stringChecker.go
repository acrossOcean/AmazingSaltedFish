package check

import "AmazingSaltedFish/utils"

type StrChecker struct {
	value string

	options []StrCheckOption
}

type StrCheckOption struct {
	operator     StrOperator
	intTarget    int
	stringTarget string
	inStrTarget  []string
}

type StrOperator int

const (
	_ StrOperator = iota
	// len(str) < ?
	StrOperatorLenLT
	// len(str) <= ?
	StrOperatorLenLE
	// len(str) = ?
	StrOperatorLenEQ
	// len(str) != ?
	StrOperatorLenNE
	// len(str) > ?
	StrOperatorLenGT
	// len(str) >= ?
	StrOperatorLenGE

	// str == ?
	StrOperatorStrEQ
	// str != ?
	StrOperatorStrNE
	// str in (?,?, ....)
	StrOperatorStrIN
)

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

func NewStrCheckOption(op StrOperator, intTarget int, strTarget string) StrCheckOption {
	var result = StrCheckOption{
		operator:     op,
		intTarget:    intTarget,
		stringTarget: strTarget,
	}

	return result
}

func NewStrCheckOptionInt(op StrOperator, intTarget int) StrCheckOption {
	var result = StrCheckOption{
		operator:  op,
		intTarget: intTarget,
	}

	return result
}

func NewStrCheckOptionStr(op StrOperator, strTarget string) StrCheckOption {
	var result = StrCheckOption{
		operator:     op,
		stringTarget: strTarget,
	}

	return result
}

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
