# `/higherorderfunctions`

Implement the CAR and CDR functions as known from Lisp, for example. Context and reference: [the native web GmbH Youtube](https://www.youtube.com/watch?v=KMdZIP9TYgQ)

Given are:

```go
type PairFunc func(int, int) int

func Cons(a, b int) func(PairFunc) int {
    return func(f PairFunc) int {
        return f(a, b)
    }
}
```

We are looking for the functions CAR and CDR, so that the following applies:

> Car(Cons(3, 4)) => 3
>
> Cdr(Cons(3, 4)) => 4

## Solution

1. We define the functions `Car` and `Cdr` as functions that each expect a function of type `PairFunc` as a parameter and return an `int`. This was virtually predefined.
2. The function `Car` should return the first number that was passed in the function `Cons`. The second number should be ignored.
3. The function `Cdr` should return the second number that was passed in the function `Cons`. The first number should be ignored.
4. We only need to call the `Cons` function in our first return statement and return the corresponding result.
