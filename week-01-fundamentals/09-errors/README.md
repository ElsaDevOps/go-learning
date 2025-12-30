## Error handling in GO
 Copy the Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

type ErrNegativeSqrt float64

and make it an error by giving it a

func (e ErrNegativeSqrt) Error() string

method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2". 

## What I learned
- **Custom error types:** You can define a type (like `type ErrNegativeSqrt float64`) and make it an error
- **Satisfying the error interface:** Any type with an `Error() string` method satisfies the `error` interface — that's the only requirement
- **Creating error instances:** `ErrNegativeSqrt(x)` converts `x` to your custom error type
- **Multiple return values:** Functions return `(result, error)` — check for error conditions first, return early if invalid
- **Format verbs:** `%f` for floats in `fmt.Sprintf`, not `%d` (that's for integers)

# Running
go run main.go