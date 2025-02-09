package util

import "golang.org/x/exp/constraints"

func PageToOffset[T constraints.Integer](page, pageSize T) (T, T) {
	if page == 1 {
		return 0, pageSize
	}
	return (page - 1) * pageSize, pageSize
}

func HasMore[T constraints.Integer](page, pageSize, count T) bool {
	return page*pageSize < count
}
