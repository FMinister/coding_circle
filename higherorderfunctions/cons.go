package higherorderfunctions

type PairFunc func(int, int) int

func COns(a, b int) func(PairFunc) int {
	return func(f PairFunc) int {
		return f(a, b)
	}
}
