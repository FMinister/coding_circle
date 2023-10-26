# `/continuousmax`

Given are a list V and an integer k. Determine the running maximum for the window size k. The runtime shall be O(n), the memory requirement shall be O(k). Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=0QPK9X1PtEE)

Given are: A simple linked list and an integer k.

Example:

> V = [27, 9, 17, 2, 12, 8]
>
> k = 3
>
> => [27, 17, 17, 12]

Conditions: The runtime shall be O(n), the memory requirement shall be O(k).

## Solution

1. First we create a slice in which we can store the result. The special thing here is that the initial length is set to 0 and the capacity to `len(values)-k+1`, i.e. the length of the given list minus the window size k plus one, to cover the maximum of all possible windows.
2. Now we create our window.
3. We go through the given list.
4. We check if the first value in the window `window` is still within our view range. If not, we remove the element.
5. Now we traverse the current window, removing all indices of values less than or equal to the current value `values[i]`.
6. We add to the window the current value at the position `i`.
7. When the window reaches the size of `k`, we add the maximum that is in the first position to our result list.
