package util

func Condition[T any](r bool, a, b T) T {
	if r {
		return a
	}
	return b
}
