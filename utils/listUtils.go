package utils

import "strings"

// IntList 提供 []int 的便利操作
type IntList []int

// RemoveDuplicate 删除 list 中相同的元素,返回删除后数组
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

// RemoveZero 删除 list 中为0 的元素, 返回删除后数组
func (receiver IntList) RemoveZero() []int {
	var result []int
	for _, i := range receiver {
		if i != 0 {
			result = append(result, i)
		}
	}

	return result
}

// StringList 提供 []string 的便利操作
type StringList []string

// RemoveDuplicate 删除 list 中相同的元素,返回删除后数组
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

// RemoveZero 删除 list 中为 长度为 0 的元素, 空格组成的字符串不会删除, 需要先调用 RemoveSpace() 后才会删除, 返回删除后数组
func (receiver StringList) RemoveZero() []string {
	var result []string
	for _, str := range receiver {
		if len(str) != 0 {
			result = append(result, str)
		}
	}

	return result
}

// RemoveSpace 删除 list 中每个 字符串的 `space`
func (receiver StringList) RemoveSpace() []string {
	var result []string
	for _, str := range receiver {
		result = append(result, strings.TrimSpace(str))
	}

	return result
}
