# `/allbutone`

Given is a list of integers. We are looking for a list with the products of all numbers except the ith digit. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=Jd_h19uabuo)

Example:

> list = [27, 9, 12, 8, 17, 2]
>
> => [29376, 88128, 66096, 99144, 46656, 396576]

Conditions: The runtime complexity should be O(n), the memory complexity should be O(n).

## Solution

Example for a small list:

```ascii
o = [1 1 1]
i = [5 2 4]

   l = 1
   o = l * o
   l = l * i
                   ┐
1. o = 1 * 1 = 1   |
1. l = 1 * 5 = 5   |
2. o = 1 * 5 = 5   ├──► o = [1 5 10]
2. l = 5 * 2 = 10  |
3. o = 5 * 2 = 10  |
                   ┘

   r = 1
   o = r * o
   r = r * i
                   ┐
1. o = 1 * 10 = 10 |
1. r = 1 * 4  = 4  |
2. o = 4 * 5  = 20 ├──► o = [10 20 8]
2. r = 4 * 2  = 8  |
3. o = 8 * 1  = 8  |
                   ┘
```

1. First we create an output array with the length of the input list. This array is initialized with 1.
2. We create an auxiliary variable `leftProduct` and set it to 1.
3. Then we iterate forwards over the input list and multiply the values at the index of the output array with the auxiliary variable `leftProduct`. We store the result at the index of the output array again.
4. We set the auxiliary variable `leftProduct` to the product of the auxiliary variable and the value of the input list at the current index.
5. ... and so on.
6. We create an auxiliary variable `rightProduct` and set it to 1.
7. Then we iterate backwards over the input list and multiply the values at the index of the output array with the auxiliary variable `rightProduct`. We save the result again at the index of the output array.
8. We set the auxiliary variable `rightProduct` to the product of the auxiliary variable and the value of the input list at the current index.
9. ... and so on.
10. We return the output array.
