# `/autocomplete`

Develop an autocomplete. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=wklh-1YTzmM)

Given are: A querystring and a list of possible terms. Now find all terms that have the querystring as a prefix.

Example:

> list = ["go", "golang", "golo", "goroutine"]
>
> q = "gol"
>
> => ["golang", "golo"]

## Solution

1. We start at the root node of the prefix tree.
2. We iterate through each character in the prefix.
3. For each character, we check if there is a child node in the `currentNode` that corresponds to the character.
4. If a child node exists for the character, we update the `currentNode` to the child node and continue with the next character in the prefix. If there is no child node for the character, we return an empty string slice to indicate that no words in the trie begin with the specified prefix.
5. If we have successfully traversed the trie using the characters in the prefix, we call the `GetWords` method on the `currentNode` with the `prefix` as argument. This method returns a list of words in the trie that have the specified prefix.
6. We return the list of words as the result of the `Find` method.
