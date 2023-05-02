package generics

import "sync"

func SumInt64(s []int64) int64 {
	sum := int64(0)

	for i := range s {
		sum += s[i]
	}

	return sum
}

func SumInt32(s []int32) int32 {
	sum := int32(0)

	for i := range s {
		sum += s[i]
	}

	return sum
}

// Interfaces can be used as type constraints! `|` is used to create type _unions_
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func SumInts[I Integer](s []I) I {
	sum := I(0)

	for i := range s {
		sum += s[i]
	}

	return sum
}

// type unions are composable
type Addable interface {
	Integer | float32 | float64 | string | complex64 | complex128
}

// Use the `~` operator to support type aliases
// type Integer interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
// }

func Sum[A Addable](s []A) A {
	var sum A

	for i := range s {
		sum += s[i]
	}

	return sum
}

func ReduceSimple[T any](f func(x, y T) T, s []T) T {
	var t T

	switch len(s) {
	case 0:
		return t
	case 1:
		return s[0]
	}

	t = f(s[0], s[1])

	if len(s) == 2 {
		return t
	}

	for _, val := range s[2:] {
		t = f(t, val)
	}

	return t
}

func Reduce[S any, T any](f func(t T, s S) T, s []S, startVal ...T) T {
	var res T

	// set initial value if provided
	if len(startVal) > 1 {
		panic("Reduce only accepts 1 starting value")
	} else if len(startVal) == 1 {
		res = startVal[0]
	}

	for i := range s {
		res = f(res, s[i])
	}

	return res
}

func Map[S any, T any](f func(e S) T, s []S) []T {
	t := make([]T, len(s))

	for i := range s {
		t[i] = f(s[i])
	}

	return t
}

func PMap[S any, T any](f func(e S) T, s []S) []T {
	t := make([]T, len(s))

	wg := &sync.WaitGroup{}

	wg.Add(len(s))
	for i := range s {
		go func(j int) {
			defer wg.Done()

			t[j] = f(s[j])
		}(i)
	}

	wg.Wait()

	return t
}
