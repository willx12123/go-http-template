package util

func PageToOffset(page, pageSize int) (int, int) {
	if page == 1 {
		return 0, pageSize
	}
	return (page - 1) * pageSize, pageSize
}

func HasMore(page, pageSize int, count int64) bool {
	return int64(page*pageSize) < count
}
