package check

type Checker interface {
	check() bool
}

func PassCheck(first Checker, more ...Checker) bool {
	if !first.check() {
		return false
	}

	for _, c := range more {
		if !c.check() {
			return false
		}
	}

	return true
}
