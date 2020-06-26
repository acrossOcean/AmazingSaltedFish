package check

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

func NewStrCheckOptionIN(strTargets []string) StrCheckOption {
	var result = StrCheckOption{
		operator:    StrOperatorStrIN,
		inStrTarget: strTargets,
	}

	return result
}

func (receiver *StrChecker) check() bool {
	for _, o := range receiver.options {
		switch o.operator {
		case StrOperatorLenLT:
			if !(len(receiver.value) < o.intTarget) {
				return false
			}
		case StrOperatorLenLE:
			if !(len(receiver.value) <= o.intTarget) {
				return false
			}
		case StrOperatorLenNE:
			if !(len(receiver.value) != o.intTarget) {
				return false
			}
		case StrOperatorLenEQ:
			if !(len(receiver.value) == o.intTarget) {
				return false
			}
		case StrOperatorLenGT:
			if !(len(receiver.value) > o.intTarget) {
				return false
			}
		case StrOperatorLenGE:
			if !(len(receiver.value) >= o.intTarget) {
				return false
			}
		case StrOperatorStrNE:
			if !(receiver.value != o.stringTarget) {
				return false
			}
		case StrOperatorStrEQ:
			if !(receiver.value == o.stringTarget) {
				return false
			}
		case StrOperatorStrIN:
			found := false
			for _, target := range o.inStrTarget {
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
