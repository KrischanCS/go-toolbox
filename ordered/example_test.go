package ordered_test

import (
	"fmt"

	"github.com/KrischanCS/go-toolbox/ordered"
)

func ExampleNew() {
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)

	fmt.Println(m.String())

	// Output:
	// (OrderedMap[{one: 1}, {two: 2}])
}

func ExampleMap_Add() {
	m := ordered.NewMap[string, int](10)
	m.Add("first", 1)
	m.Add("second", 2)
	m.Add("third", 3)

	for k, v := range m.All() {
		fmt.Printf("%s: %d\n", k, v)
	}

	// Output:
	// first: 1
	// second: 2
	// third: 3
}

func ExampleMap_Get() {
	m := ordered.NewMap[string, int](10)
	m.Add("key", 42)

	value, ok := m.Get("key")
	fmt.Printf("value: %d, ok: %t\n", value, ok)

	value, ok = m.Get("nonexistent")
	fmt.Printf("value: %d, ok: %t\n", value, ok)

	// Output:
	// value: 42, ok: true
	// value: 0, ok: false
}

func ExampleMap_Delete() {
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	deleted, ok := m.Delete("two")
	fmt.Printf("deleted: %d, ok: %t\n", deleted, ok)
	fmt.Printf("len: %d\n", m.Len())

	for k, v := range m.All() {
		fmt.Printf("%s: %d\n", k, v)
	}

	// Output:
	// deleted: 2, ok: true
	// len: 2
	// one: 1
	// three: 3
}

func ExampleMap_Contains() {
	m := ordered.NewMap[string, int](10)
	m.Add("key", 42)

	fmt.Println(m.Contains("key"))
	fmt.Println(m.Contains("other"))

	// Output:
	// true
	// false
}

func ExampleMap_Len() {
	m := ordered.NewMap[string, int](10)
	fmt.Println(m.Len())

	m.Add("one", 1)
	m.Add("two", 2)
	fmt.Println(m.Len())

	// Output:
	// 0
	// 2
}

func ExampleMap_All() {
	m := ordered.NewMap[string, int](10)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)

	// Iteration always returns elements in insertion order
	for k, v := range m.All() {
		fmt.Printf("%s: %d\n", k, v)
	}

	// Output:
	// a: 1
	// b: 2
	// c: 3
}

func ExampleNewSet() {
	s := ordered.NewSet[string](10)
	s.Add("one")
	s.Add("two")

	fmt.Println(s.Len())

	// Output:
	// 2
}

func ExampleSetOf() {
	s := ordered.SetOf("one", "two", "three")

	for v := range s.All() {
		fmt.Println(v)
	}

	// Output:
	// one
	// two
	// three
}

func ExampleSet_Add() {
	s := ordered.NewSet[string](10)
	s.Add("first")
	s.Add("second")
	s.Add("third")

	for v := range s.All() {
		fmt.Println(v)
	}

	// Output:
	// first
	// second
	// third
}

func ExampleSet_Remove() {
	s := ordered.SetOf("one", "two", "three")

	ok := s.Remove("two")
	fmt.Printf("removed: %t\n", ok)
	fmt.Printf("len: %d\n", s.Len())

	for v := range s.All() {
		fmt.Println(v)
	}

	// Output:
	// removed: true
	// len: 2
	// one
	// three
}

func ExampleSet_Contains() {
	s := ordered.SetOf("one", "two", "three")

	fmt.Println(s.Contains("two"))
	fmt.Println(s.Contains("four"))

	// Output:
	// true
	// false
}

func ExampleSet_Values() {
	s := ordered.SetOf("a", "b", "c")

	// Values returns a slice in insertion order
	values := s.Values()
	fmt.Println(values)

	// Output:
	// [a b c]
}

func ExampleSet_All() {
	s := ordered.SetOf("a", "b", "c")

	// Iteration always returns elements in insertion order
	for v := range s.All() {
		fmt.Println(v)
	}

	// Output:
	// a
	// b
	// c
}

func ExampleSet_Clone() {
	s := ordered.SetOf("one", "two", "three")
	clone := s.Clone()

	// Modifying original doesn't affect clone
	s.Add("four")

	fmt.Printf("original len: %d\n", s.Len())
	fmt.Printf("clone len: %d\n", clone.Len())

	// Output:
	// original len: 4
	// clone len: 3
}

func ExampleSet_String() {
	s := ordered.SetOf("one", "two", "three")
	fmt.Println(s.String())

	empty := ordered.NewSet[string](10)
	fmt.Println(empty.String())

	// Output:
	// (OrderedSet[string]: [one, two, three])
	// (OrderedSet[string]: <empty>)
}
