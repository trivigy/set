package set

import (
	"fmt"
	"sort"
)

func ExampleNew() {
	m := New("a", "a", "b")
	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a b]
}

func ExampleSet_Add() {
	m := New("a", "a", "a")
	m.Add("b", "b", "c", "d")
	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a b c d]
}

func ExampleSet_Contains() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.Contains("b", "c", "d"))
	// Output:
	// true
}

func ExampleSet_NotContains() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.Contains("f", "d", "g"))
	// Output:
	// false
}

func ExampleSet_Equals() {
	m1 := New("b", "b", "c", "d")
	m2 := New("c", "b", "d", "b")
	fmt.Println(m1.Equals(m2))
	// Output:
	// true
}

func ExampleSet_NotEquals() {
	m1 := New("b", "b", "c", "d", "f")
	m2 := New("c", "b", "d", "b", "g")
	fmt.Println(m1.Equals(m2))
	// Output:
	// false
}

func ExampleSet_Clear() {
	m := New("b", "b", "c", "d")
	m.Clear()
	fmt.Println(m)
	// Output:
	// []
}

func ExampleSet_IsEmpty() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.IsEmpty())
	// Output:
	// false
}

func ExampleSet_Size() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.Size())
	// Output:
	// 3
}

func ExampleSet_Remove() {
	m := New("a", "a", "a", "b", "b")
	m.Remove("a", "a")
	fmt.Println(m)
	// Output:
	// [b]
}

func ExampleSet_Iter() {
	m := New("a", "a", "b")
	for elem := range m.Iter() {
		fmt.Println(elem)
	}
	// Output:
	// a
	// b
}

func ExampleSet_Subset() {
	m1 := New("a", "b", "c")
	m2 := New("c", "b")
	fmt.Println(m2.Subset(m1))
	// Output:
	// true
}

func ExampleSet_Superset() {
	m1 := New("a", "b", "c")
	m2 := New("c", "b")
	fmt.Println(m1.Superset(m2))
	// Output:
	// true
}

func ExampleSet_Union() {
	m1 := New("a", "b", "c")
	m2 := New("c", "d", "e")
	m := m1.Union(m2)

	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a b c d e]
}

func ExampleSet_Intersect() {
	m1 := New("a", "b", "c")
	m2 := New("c", "d", "e")
	m := m1.Intersect(m2)

	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [c]
}

func ExampleSet_Diff() {
	m1 := New("a", "b", "c")
	m2 := New("c", "d", "e")
	m := m1.Diff(m2)

	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a b]
}

func ExampleSet_SymDiff() {
	m1 := New("a", "b", "c")
	m2 := New("c", "d", "e")
	m := m1.SymDiff(m2)

	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a b d e]
}
