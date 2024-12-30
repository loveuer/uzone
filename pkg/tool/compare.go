package tool

func Min[T string | int64 | uint64 | int | uint](a, b T) T {
	if a < b {
		return a
	}

	return b
}
