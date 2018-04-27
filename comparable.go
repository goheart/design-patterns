// Copyright 2018 zjj2wry. All rights reserved.

// This package contains various uses of the go interface,
// which can help me achieve a variety of design patterns

package comparable

// Comparable interface ComparaTo function parameters and return values ​​are also Comparable interface.
type Comparable interface {
	CompareTo(to Comparable) Comparable
}

// guarantee Comparator implement Comparable interface.
var _ Comparable = new(Comparator)

// Comparator combinate and implement Comparable interface
type Comparator struct {
	Comparable Comparable
}

// New return Comparator.
func New(c Comparable) Comparator {
	return Comparator{
		c,
	}
}

// Any func can convert to Comparefun. eg: compareFun(o.Any), so it alse implement
// Comparable interface.
func (o Comparator) Any(to Comparable) Comparable {
	return to.CompareTo(o)
}

// Get return comparatable.
// Comparator implement Comparable interface. So it can alse as
// instance return.
func (o Comparator) Get() Comparable {
	return o
}

// CompareTo is CompareFunc type and alse help Comparator implement Comparable interface.
func (o Comparator) CompareTo(to Comparable) Comparable {
	return o.Comparable.CompareTo(to)
}

// CompareFunc implement the Comparable interface for func
type CompareFunc func(to Comparable) Comparable

// CompareTo return func(to)
func (o CompareFunc) CompareTo(to Comparable) Comparable {
	return o(to)
}

// Int implements the Comparable interface for integers.
type Int int

// CompareTo returns int(a) - int(b).
func (a Int) CompareTo(b Comparable) Comparable {
	if (int)(a-b.(Int)) > 0 {
		return a
	}
	return b
}
