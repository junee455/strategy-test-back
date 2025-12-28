package helpers

func Includes[T comparable](in []T, element T) bool {
	for _, v := range in {
		if v == element {
			return true
		}
	}
	return false
}
