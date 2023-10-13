# `/median`

A function is searched for that calculates the continuously updated median for a list of numbers and returns these values as a list. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=28ad9BYDHlI)

Example:

> [17, 2, 8, 27, 12, 9]
>
> => [17, 9.5, 8, 12.5, 12, 10.5]

## Solution

The approach is as follows:

- First calculate if the list length is even or odd.
- Then the list is sorted.
- Depending on whether the list is even or odd, the median is calculated differently:
  - Even: The two middle numbers are added and divided by 2.
  - Odd: The middle number is returned.
