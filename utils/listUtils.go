package utils

import "strings"

type IntList []int

func (receiver IntList) RemoveDuplicate() []int {
	var result []int
	var m = make(map[int]struct{})
	for _, i := range receiver {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
			result = append(result, i)
		}
	}

	return result
}

func (receiver IntList) RemoveZero() []int {
	var result []int
	for _, i := range result {
		if i != 0 {
			result = append(result, i)
		}
	}

	return result
}

type StringList []string

func (receiver StringList) RemoveDuplicate() []string {
	var result []string
	var m = make(map[string]struct{})
	for _, str := range receiver {
		if _, ok := m[str]; !ok {
			m[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func (receiver StringList) RemoveZero() []string {
	var result []string
	for _, str := range result {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}

func (receiver StringList) RemoveSpace() []string {
	var result []string
	for _, str := range result {
		result = append(result, strings.TrimSpace(str))
	}

	return result
}
