# `/remove_the_last_element`

Remove the i-th element from the end of a list. Context and reference:[the native web GmbH Youtube](https://www.youtube.com/watch?v=kPR3IWC0Wvc)

Given are: A simple linked list and an integer i.

Example:

> [1, 2, 3, 4, 5]
>
> i = 2
>
> => [1, 2, 3, 5]

## Solution

The approach is as follows:

1. Two pointers, `lead` and `follow`, are initially set to the head of the linked list. These pointers are used to traverse the list while maintaining a gap of `k` nodes between them.

2. A `for` loop is used to move the `lead` pointer `k` positions ahead in the list. If `lead` becomes `nil` (i.e., the list has fewer than `k` nodes), it means there's no `k`-th last element to remove, so the function returns `nil` (indicating an invalid operation).

3. After the first loop, if `lead` is still not `nil`, it means there are at least `k` nodes in the list. Now, both `lead` and `follow` start moving through the list, one node at a time, until `lead` reaches the end of the list.

4. When `lead` reaches the end (i.e., `lead.Next` becomes `nil`), the `follow` pointer is positioned at the node just before the `k`-th last node. This is because `follow` is `k` nodes behind `lead`.

5. The last element is removed by updating the `follow.Next` pointer to skip the `k`-th last node. This is done by setting `follow.Next` to `follow.Next.Next`, effectively bypassing the `k`-th last node.

6. The function returns the original `head` of the linked list, which may have been updated if the `k`-th last element was the first element in the list. If the first element is removed, `head` is updated to point to the second element.
