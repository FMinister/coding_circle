# `/warpspeedsorting`

Sort a large number of postal codes as quickly as possible. A function that receives a list of zip codes, sorts them and returns the sorted list. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=g6-DvmV4OYI)

Conditions: The runtime complexity should be better than O(n log n), ideally O(n). Ideally, the memory requirement should also be O(n).

## Solution

1. First we take out the special cases, namely if none or only one element is contained in the list. Here you can simply return the list.
2. Then we create an auxiliary function that determines the maximum and minimum of the list. We simply iterate over the list and compare each element with the current minimum and maximum. The minimum and maximum are then returned.
3. With the minimum and maximum we now create a `countingArray`, which can have the size of maximum minus minimum plus 1, since the range of all given zip codes will fit in there and initialize it with 0 at each position.
4. Then we go through the list of the given zip codes and increase the value in the `countingArray` at the position `postcode - minimum` by 1 and have thus counted the occurrences of all zip codes.
5. Now we create the array for the sorted postal codes with a length of zero and a maximum length of the length of the given list of postal codes.
6. Then we iterate over the `countingArray`, take the current index and the number of postal codes and insert the postal code into the sorted array as often as it occurs in the `countingArray`. To do this, we append the zip code, which is calculated with `index + minimum`, to the sorted array. We do this as often as the zip code occurs in the `countingArray`, which is why we iterate over the number of zip codes in the `countingArray` in this loop and reduce it by one after appending.
7. At the end we return the sorted array.
