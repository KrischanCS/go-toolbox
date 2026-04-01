package stack_test

import (
	"fmt"

	"github.com/KrischanCS/go-toolbox/stack"
)

func ExampleOf() {
	s := stack.Of(1, 2, 3)

	fmt.Println(s.String())

	// Output:
	// (Stack[int]: Top< 3, 2, 1 |Bottom)
}

func ExampleStack_Push() {
	s := stack.Of[int]()

	fmt.Println("Initial: ", s.String())

	s.Push(1)
	fmt.Printf("Push(1): %s\n", s.String())

	s.Push(2)
	fmt.Printf("Push(2): %s\n", s.String())

	// Output:
	//
	// Initial:  (Stack[int]: <empty>)
	// Push(1): (Stack[int]: Top< 1 |Bottom)
	// Push(2): (Stack[int]: Top< 2, 1 |Bottom)
}

func ExampleStack_Pop() {
	s := stack.Of(1, 2)

	fmt.Println("Initial: ", s.String())
	fmt.Println()

	value, ok := s.Pop()
	fmt.Printf("Pop(): %d, %t\n", value, ok)
	fmt.Printf("Remaining: %s\n", s.String())
	fmt.Println()

	value, ok = s.Pop()
	fmt.Printf("Pop(): %d, %t\n", value, ok)
	fmt.Printf("Remaining: %s\n", s.String())
	fmt.Println()

	value, ok = s.Pop()
	fmt.Printf("Pop(): %d, %t\n", value, ok)

	// Output:
	//
	// Initial:  (Stack[int]: Top< 2, 1 |Bottom)
	//
	// Pop(): 2, true
	// Remaining: (Stack[int]: Top< 1 |Bottom)
	//
	// Pop(): 1, true
	// Remaining: (Stack[int]: <empty>)
	//
	// Pop(): 0, false
}

func ExampleStack_PopN() {
	s := stack.Of(1, 2, 3, 4)

	fmt.Println("Initial: ", s.String())
	fmt.Println()

	values := s.PopN(2)
	fmt.Printf("PopN(2): %v\n", values)
	fmt.Printf("Remaining: %s\n", s.String())
	fmt.Println()

	values = s.PopN(5)
	fmt.Printf("PopN(5): %v\n", values)
	fmt.Printf("Remaining: %s\n", s.String())

	// Output:
	//
	// Initial:  (Stack[int]: Top< 4, 3, 2, 1 |Bottom)
	//
	// PopN(2): [4 3]
	// Remaining: (Stack[int]: Top< 2, 1 |Bottom)
	//
	// PopN(5): [2 1]
	// Remaining: (Stack[int]: <empty>)
}

func ExampleStack_PopAll() {
	s := stack.Of(1, 2, 3)

	fmt.Println("Initial: ", s.String())
	fmt.Println()

	values := s.PopAll()
	fmt.Printf("PopAll(): %v\n", values)
	fmt.Printf("Remaining: %s\n", s.String())

	// Output:
	//
	// Initial:  (Stack[int]: Top< 3, 2, 1 |Bottom)
	//
	// PopAll(): [3 2 1]
	// Remaining: (Stack[int]: <empty>)
}

func ExampleStack_Peek() {
	s := stack.Of(1, 2)

	fmt.Println("Initial: ", s.String())
	fmt.Println()

	value, ok := s.Peek()
	fmt.Printf("Peek(): %d, %t\n", value, ok)
	fmt.Printf("Remaining: %s\n", s.String())

	// Output:
	//
	// Initial:  (Stack[int]: Top< 2, 1 |Bottom)
	//
	// Peek(): 2, true
	// Remaining: (Stack[int]: Top< 2, 1 |Bottom)
}

func ExampleStack_PeekN() {
	s := stack.Of(1, 2, 3, 4)

	fmt.Println("Initial: ", s.String())
	fmt.Println()

	values := s.PeekN(2)
	fmt.Printf("PeekN(2): %v\n", values)
	fmt.Printf("Remaining: %s\n", s.String())

	// Output:
	//
	// Initial:  (Stack[int]: Top< 4, 3, 2, 1 |Bottom)
	//
	// PeekN(2): [4 3]
	// Remaining: (Stack[int]: Top< 4, 3, 2, 1 |Bottom)
}

func ExampleStack_PeekAll() {
	s := stack.Of(1, 2, 3)

	fmt.Println("Initial: ", s.String())
	fmt.Println()

	values := s.PeekAll()
	fmt.Printf("PeekAll(): %v\n", values)
	fmt.Printf("Remaining: %s\n", s.String())

	// Output:
	//
	// Initial:  (Stack[int]: Top< 3, 2, 1 |Bottom)
	//
	// PeekAll(): [3 2 1]
	// Remaining: (Stack[int]: Top< 3, 2, 1 |Bottom)
}

func ExampleStack_String() {
	s := stack.Of(1, 2, 3)

	fmt.Println(s.String())

	// Output:
	// (Stack[int]: Top< 3, 2, 1 |Bottom)
}
