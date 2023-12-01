package higherorderfunctions

type PairFunc func(int, int) int

func Cons(a, b int) func(PairFunc) int {
	return func(f PairFunc) int {
		return f(a, b)
	}
}
