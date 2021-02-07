// Package set implements the set data structure and the fundamental operations on sets.
package set

import (
	"errors"
	"fmt"
)

// Set type models sets.
// The elements of a set are the keys of the underlying map.
// Since a set is a map, its elements (the keys of the map) must be comparable types,
// that is boolean, numeric, string, pointer, channel or interface types,
// or structs or arrays that contain only those types.
type Set map[interface{}]struct{}

// Pair type models ordered pairs (for the cartesian product).
type Pair [2]interface{}

// New creates a new set with the specified elements and returns a pointer to it.
func New(elems ...interface{}) *Set {
	a := make(Set)
	for _, e := range elems {
		a.Add(e)
	}
	return &a
}

// Add adds an element to a set.
func (set *Set) Add(elem interface{}) {
	(*set)[elem] = struct{}{}
}

// Remove removes an element from a set.
// If the element is not in the set, it's a no-op.
func (set *Set) Remove(elem interface{}) {
	delete(*set, elem)
}

// Has checks if a set has the specified element.
func (set *Set) Has(elem interface{}) bool {
	_, ok := (*set)[elem]
	return ok
}

// IsSubsetOf checks if a set is a subset of the specified set.
func (set *Set) IsSubsetOf(set2 *Set) bool {
	for e := range *set {
		if !set2.Has(e) {
			return false
		}
	}
	return true
}

// IsSupersetOf checks if a set is a superset of the specified set.
func (set *Set) IsSupersetOf(set2 *Set) bool {
	return set2.IsSubsetOf(set)
}

// IsEmpty checks if a set is the empty set.
func (set *Set) IsEmpty() bool {
	return set.Card() == 0
}

// GetOne returns a random element from a set and an error.
// If the set is empty it returns nil and the error "Set is empty".
func (set *Set) GetOne() (interface{}, error) {
	for e := range *set {
		return e, nil
	}
	return nil, errors.New("Set is empty")
}

// Card returns the cardinal of a set.
func (set *Set) Card() int {
	return len(*set)
}

// Equal checks if two sets are equal.
func Equal(setA, setB *Set) bool {
	return setA.IsSubsetOf(setB) && setA.IsSupersetOf(setB)
}

// Print prints a set in a format similar to that of struct.
func Print(set *Set) {
	fmt.Printf("{ ")
	for e := range *set {
		switch e.(type) {
		case rune, string:
			fmt.Printf("%q ", e)
		default:
			fmt.Printf("%v ", e)
		}
	}
	fmt.Printf("}\n")
}

// Union returns the union of two sets.
func Union(setA, setB *Set) *Set {
	s := make(Set)
	for e := range *setA {
		s.Add(e)
	}
	for e := range *setB {
		s.Add(e)
	}
	return &s
}

// Inter returns the intersection of two sets.
func Inter(setA, setB *Set) *Set {
	s := make(Set)
	for e := range *setA {
		if setB.Has(e) {
			s.Add(e)
		}
	}
	return &s
}

// Diff returns the set difference setA\setB.
func Diff(setA, setB *Set) *Set {
	s := make(Set)
	for e := range *setA {
		if !setB.Has(e) {
			s.Add(e)
		}
	}
	return &s
}

// SymDiff returns the symmetric difference of two sets.
func SymDiff(setA, setB *Set) *Set {
	s := make(Set)
	for e := range *setA {
		s.Add(e)
	}
	for e := range *setB {
		if setA.Has(e) {
			s.Remove(e)
		} else {
			s.Add(e)
		}
	}
	return &s
}

// Prod returns the cartesian product of two sets.
func Prod(setA, setB *Set) *Set {
	s := make(Set)
	for e1 := range *setA {
		for e2 := range *setB {
			s.Add(Pair{e1, e2})
		}
	}
	return &s
}
