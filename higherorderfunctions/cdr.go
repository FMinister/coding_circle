package higherorderfunctions

func Cdr(cons func(pair PairFunc) int) int {
	return cons(func(a, b int) int {
		return b
	})
}
