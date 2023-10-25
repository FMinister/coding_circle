package tictactoe

type Coordinate struct {
	Row int
	Col int
}

func HasWinningPosition(n int, coordinates []Coordinate) bool {
	rowCounter := make([]int, n)
	colCounter := make([]int, n)
	diagonalLeftToRightCounter, diagonalRightToLeftCounter := 0, 0

	for _, coordinate := range coordinates {
		row, col := coordinate.Row, coordinate.Col

		rowCounter[row]++
		colCounter[col]++

		if row == col {
			diagonalLeftToRightCounter++
		}

		if row+col == n-1 {
			diagonalRightToLeftCounter++
		}

		if rowCounter[row] == n || colCounter[col] == n || diagonalLeftToRightCounter == n || diagonalRightToLeftCounter == n {
			return true
		}
	}

	return false
}
