package util

func GreatestInt(ints ...int) int {
	var greatest int
	for _, i := range ints {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}
