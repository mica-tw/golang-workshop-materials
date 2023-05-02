# generics-exercises

This repo contains some demonstrations of generic types and functions and how to use them.

Code intended for demo purposes only, not production-ready.

## How to use

Start by commenting out most of the content in `demo/main.go`.

Uncomment each line starting with `fmt.Printf(...` one-by-one, along with any code preceding each
statement. Make sure you understand the concepts introduced in each example before moving on to 
the next one.

Some lines don't compile. If there's a comment mentioning that it won't compile,
read the error from the compiler and try to understand it. Otherwise, make the
code changes necessary for it to compile and keep going. 

### Memory Store Exercise

The last part of `demo/main.go` shows the intended behaviour of the `mem_store` package. To finish,
implement the missing functionality in `pkg/mem_store/mem_store.go`. Method stubs are provided for
guidance, but feel free to add functions, refactor, or change anything necessary.

## Standard Library Generics

Since generics are relatively new, very few std lib packages make use of generics. 

However, some experimental packages that might later be in the standard library use generics. 
These packages may break in between _minor_ Go versions, although they tend to be stable enough
for production in well-tested code bases.

- https://pkg.go.dev/golang.org/x/exp/slices
- https://pkg.go.dev/golang.org/x/exp/maps
- https://pkg.go.dev/golang.org/x/exp/constraints

## Notes and Further reading

If you're not satisfied already, these two blog posts would be a good next step for reading:

- https://go.dev/blog/intro-generics
- https://tip.golang.org/blog/when-generics

And finally a note for the extreme edge cases. The predeclared type `comparable` is not _always_ safe to
use. If you want to get into the academic theory and when the type system behaves in unintuitive ways,
read this: https://go.dev/blog/comparable 
