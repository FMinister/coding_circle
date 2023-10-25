# `/tictactoe`

Given are the playing field length n and a list of coordinates. We are looking for a function which decides for a playing field of size n*n whether the coordinates contain a winning position. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=ofs5v4jXa-k)

Example:

> n = 3
>
> coordinates = [(0, 0), (0, 2), (1, 1), (2, 2)]
>
> => true

Conditions: Runtime O(n^2), Memory O(n), Single-pass.

## Solution

The approach is as follows:

1. We create a list of length n for row and column, where each element represents a row or column. We also create two count variables for the diagonals.
2. We iterate over the list of coordinates and take the row and column values. For each corresponding row and column the value increases by 1.
3. We check if the value of the row is equal to the value of the column. If it is, we increase the value of the diagonal running from the upper left to the lower right by 1. Similarly, we check the case when the sum of the row and column coordinates is equal to n-1. If this is the case, we increase the value of the diagonal running from the upper right to the lower left by 1.
4. Finally we check if one of the created counters is equal to the value of n. If it is, we return true.
5. If we have run through all values of the coordinates list and none of the created counters is equal to the value of n, we return a false.
