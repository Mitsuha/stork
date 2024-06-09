package slices

func Container[T comparable](in []T, el T) bool {
	for _, t := range in {
		if t == el {
			return true
		}
	}
	return false
}
