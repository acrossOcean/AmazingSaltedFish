package check

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
	for _, o := range receiver.options {
		switch o.operator {
		case IntOperatorLT:
			if !(receiver.value < o.target) {
				return false
			}
		case IntOperatorLE:
			if !(receiver.value <= o.target) {
				return false
			}
		case IntOperatorNE:
			if !(receiver.value != o.target) {
				return false
			}
		case IntOperatorEQ:
			if !(receiver.value == o.target) {
				return false
			}
		case IntOperatorGT:
			if !(receiver.value > o.target) {
				return false
			}
		case IntOperatorGE:
			if !(receiver.value >= o.target) {
				return false
			}
		case IntOperatorIN:
			found := false
			for _, target := range o.inTarget {
				if receiver.value == target {
					found = true
					break
				}
			}
			return found
		default:
			return false
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
