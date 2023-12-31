# `/floydscirclefinding`

Given is a single-linked list. We are looking for an algorithm that decides whether the list contains a loop. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=cH3WLZn3QmA)

Example:

```ascii
12 ─► 8 ─► 17 ─► 2 ─► 27 ─► 9
           ▲                │
           └────────────────┘

=> true
```

Conditions: The runtime complexity should be O(n), the memory complexity should be O(1).

## Solution

1. We create two pointers, one slow and one fast.
2. The slow one iterates through the list one element at a time.
3. The fast one iterates through the list in two steps.
4. If the list has a loop, the fast pointer will eventually be equal to the slow pointer, we return a `true`.
5. If the list has no loop, the fast pointer will eventually reach the end of the list, we return a `false`.
