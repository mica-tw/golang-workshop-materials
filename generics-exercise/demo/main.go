package main

import (
	"fmt"
	"generics-exercise/pkg/generics"
	"generics-exercise/pkg/util"
	"reflect"
	"sort"

	"golang.org/x/exp/slices"
)

const fooConst = "foo"

func main() {
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/////////// Sorting                      //////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// How did sorting use to work in Go?
	toSort := []string{"my", "slice", "of", "strings", "to", "be", "sorted"}
	toSortCopy := make([]string, len(toSort))
	copy(toSortCopy, toSort)

	sort.Strings(toSort) // individual implementations for every type 🤦‍♀️
	fmt.Printf("sort(%v) = %v\n", toSortCopy, toSort)

	// But what if I wanted to sort them in a non-standard order? E.g. short-lex instead of alphabetical
	copy(toSortCopy, toSort)
	//sorting functions have to address slice elements by index. Not memory safe at all 👎
	sort.Slice(toSort, func(i, j int) bool {
		if len(toSort[i]) == len(toSort[j]) {
			return toSort[i] < toSort[j]
		}

		return len(toSort[i]) < len(toSort[j])
	})
	fmt.Printf("sortShortLex(%v) = %v\n", toSortCopy, toSort)

	// Generics remove this problem
	toSort = []string{"my", "slice", "of", "strings", "to", "be", "sorted"}
	copy(toSortCopy, toSort)

	// can be used on any type that satisfies `constraints.Ordered`, i.e. any type that can use `<`
	slices.Sort(toSort)
	fmt.Printf("with generics: sort(%v) = %v\n", toSortCopy, toSort)

	// and with a custom ordering, we have memory safety
	copy(toSortCopy, toSort)
	slices.SortFunc(toSort, func(x, y string) bool {
		if len(x) == len(y) {
			return x < y
		}

		return len(x) < len(y)
	})
	fmt.Printf("with generics: sortShortLex(%v) = %v\n", toSortCopy, toSort)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/////////// Pointer Shenanigans                       /////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// How do you get a pointer if you don't already have a variable?

	// Can't even dereference string without variable
	// fmt.Printf("typeOf(%v) = %v\n", &"foo", reflect.TypeOf(&"foo"))

	// can't dereference a constant either... 🤔
	// fmt.Printf("typeOf(%v) = %v\n", &fooConst, reflect.TypeOf(&fooConst))

	// you can put it in a variable, but that's annoying
	foo := "foo"
	fmt.Printf("typeOf(%v) = %v\n", &foo, reflect.TypeOf(&foo))

	// you can write a function to get a string pointer
	fmt.Printf("typeOf(%v) = %v\n", util.GetStrPtr("foo"), reflect.TypeOf(util.GetStrPtr("foo")))

	// But what if I want to get a pointer for an int now? Writing a new function for every type will get tedious fast
	// Generics to the rescue!
	fmt.Printf("typeOf(%v) = %v\n", util.GetPtr("foo"), reflect.TypeOf(util.GetPtr("foo")))
	fmt.Printf("typeOf(%v) = %v\n", util.GetPtr(3), reflect.TypeOf(util.GetPtr(3)))

	// the same function can get a pointer for any value.
	// it can even get a pointer pointer pointer... 🤯
	fmt.Printf("typeOf(%v) = %v\n", util.GetPtr(util.GetPtr("🤯")), reflect.TypeOf(util.GetPtr(util.GetPtr("🤯"))))
	fmt.Printf("typeOf(%v) = %v\n", util.GetPtr(util.GetPtr(util.GetPtr("🤯"))), reflect.TypeOf(util.GetPtr(util.GetPtr(util.GetPtr("🤯")))))

	// even nil can be ptr'd
	fmt.Printf("typeOf(%v) = %v\n", util.GetPtr[any](nil), reflect.TypeOf(util.GetPtr[any](nil)))

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/////////// Summation and Reduction                       /////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	ints64 := []int64{3, 2, 5, 8, 10, -3, 0}
	ints32 := []int32{3, 2, 5, 8, 10, -3, 0}

	// without generics we have to have separate sum methods for each type
	fmt.Printf("sum(%v) = %d\n", ints64, generics.SumInt64(ints64))
	// fmt.Printf("sum(%v) = %d\n", ints32, generics.SumInt64(ints32)) // this doesn't compile
	fmt.Printf("sum(%v) = %d\n", ints32, generics.SumInt32(ints32))

	// but we can define a generic integer type
	fmt.Printf("sum(%v) = %v = sum(%v) = %v\n", ints64, generics.SumInts(ints64), ints32, generics.SumInts(ints32))

	// What about using it for floats though?
	// Or to concatenate strings?
	floats := []float64{3, 2, 5, 8, 10, -3, 0}
	strings := []string{"three", "two", "five", "eight", "ten", "negative three", "zero"}
	fmt.Printf("sum(%v) = %v\n", floats, generics.Sum(floats))
	fmt.Printf("sum(%v) = %v\n", strings, generics.Sum(strings))

	// But what if I use custom types?
	// type myInt int
	// myInts := []myInt{3, 2, 5, 8, 10, -3, 0}
	// How can I make this work? Unioning every single possible type is getting tiring...
	// fmt.Printf("sum(%v) = %v\n", myInts, generics.Sum(myInts))

	// Okay that's cool, what if I use types that don't support the `+` operator though?
	intPtrs := []*int{util.GetPtr(3), util.GetPtr(2), util.GetPtr(5), util.GetPtr(8), util.GetPtr(10), util.GetPtr(-3), util.GetPtr(0)}
	fmt.Printf("sum(%v) = %v\n", intPtrs, *generics.ReduceSimple(
		func(x, y *int) *int {
			return util.GetPtr(*x + *y)
		},
		intPtrs,
	))

	// Now we can also use this to do weird calculations
	fmt.Printf("or for more complicated operations: weird_calculation(%v)=%f\n", floats, generics.ReduceSimple(
		func(x float64, y float64) float64 {
			return x + (y / 2)
		},
		floats,
	))

	// map
	fmt.Printf("But that last one could be better done by mapping: even_odd(%v)=%v\n", ints64, generics.Map(
		func(n int64) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		},
		ints64,
	))

	fmt.Printf("Or use it to abstract concurrency: even_odd(%v)=%v\n", ints64, generics.PMap(
		func(n int64) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		},
		ints64,
	))

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/////////// You can make generic structs & interfaces too                   ///////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// Uncomment the code below, and complete the mem_store package to behave as expected.
	// hint: the only code you have to touch is in pkg/mem_store/mem_store.go

	// Create a store for some `product`
	// prdStore := mem_store.NewMemStore[string, products.Product]()
	// prd, _ := prdStore.Store("ad-id", products.Product{
	// 	ID:          "prd-id",
	// 	Title:       "my product title",
	// 	Price:       32,
	// 	Description: "my product description",
	// })
	// foundPrd, _ := prdStore.Find(prd.ID)
	// fmt.Printf("We can use this memory store just like a non-generic one: find(%s) = %+v\n", prd.ID, *foundPrd)

	// // it can also store other things
	// stringStore := mem_store.NewMemStore[string, string]()
	// stringStore.Store("string-id", "string-value")
	// foundStr, _ := stringStore.Find("string-id")
	// fmt.Printf("This store is very versatile: find(%s) = %s\n", "string-id", *foundStr)

	// // but the prdStore can't store non-products!
	// // prdStore.Store("foo-id", "foo-string") // Does not compile

	// // How about validation?
	// validatorPrdStore := mem_store.NewMemStoreWithValidation[string](func(prd products.Product) error {
	// 	if len(prd.Description) > 20 {
	// 		return fmt.Errorf("description must be 20 characters or less")
	// 	}

	// 	return nil
	// })

	// prdToStore := products.Product{
	// 	Title:       "product title",
	// 	Description: "more than twenty character long",
	// 	ID:          "prd-id",
	// 	Price:       32,
	// }
	// prd, err := validatorPrdStore.Store("prd-id", prdToStore)
	// fmt.Printf("Store(%+v) = prd(%+v), error(%v)\n", prdToStore, prd, err)
}
