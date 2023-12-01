package higherorderfunctions

func Car(cons func(pair PairFunc) int) int {
	return cons(func(a, b int) int {
		return a
	})
}
