package check

type checker interface {
	check() bool
}

// PassCheck 检查所有检查器是否都通过了
func PassCheck(first checker, more ...checker) bool {
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
