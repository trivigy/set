// Package set implements a map-based set data structure.
package set

import (
	"fmt"
	"sync"
)

// Set represents the set data structure.
type Set struct {
	sync.RWMutex
	elems map[interface{}]struct{}
}

// New creates and returns a reference to an empty set. Operations on the
// resulting set are thread-safe.
func New(l ...interface{}) *Set {
	s := Set{elems: make(map[interface{}]struct{})}
	for _, e := range l {
		s.add(e)
	}
	return &s
}

// String returns a string representation of the set.
func (s *Set) String() string {
	return fmt.Sprint(s.ToSlice())
}

// Add adds occurrences of each of the specified elements to this set.
func (s *Set) Add(l ...interface{}) {
	s.Lock()
	defer s.Unlock()
	for _, e := range l {
		s.add(e)
	}
}

func (s *Set) add(e interface{}) bool {
	if _, found := s.elems[e]; found {
		return false
	}

	s.elems[e] = struct{}{}
	return true
}

// Contains determines whether this set contains the occurrence of each of the
// specified elements. If at least one occurrence missing returns false;
// otherwise true.
func (s *Set) Contains(l ...interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	return s.contains(l...)
}

func (s *Set) contains(l ...interface{}) bool {
	for _, e := range l {
		if _, ok := s.elems[e]; !ok {
			return false
		}
	}
	return true
}

// Equals compares the specified set with this set for equality. Returns true if
// this set contains equal elements; false otherwise.
func (s *Set) Equals(o *Set) bool {
	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()
	return s.equals(o)
}

func (s *Set) equals(o *Set) bool {
	if len(s.elems) != len(o.elems) {
		return false
	}

	for e := range s.elems {
		if !o.contains(e) {
			return false
		}
	}
	return true
}

// Clear removes all of the elements from this set. The collection will be empty
// after this method returns.
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.elems = make(map[interface{}]struct{})
}

// IsEmpty returns true if this set contains no elements; false otherwise.
func (s *Set) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.size() == 0
}

// Size returns the total number of elements in this multiset.
func (s *Set) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.size()
}

func (s *Set) size() int {
	return len(s.elems)
}

// Remove removes each of the specified elements from this set, if present.
// Returns true if this set changed as a result of the call
func (s *Set) Remove(l ...interface{}) bool {
	s.Lock()
	defer s.Unlock()
	var changed bool
	for _, e := range l {
		if count := s.remove(e); count {
			changed = true
		}
	}
	return changed
}

func (s *Set) remove(e interface{}) bool {
	if _, found := s.elems[e]; found {
		delete(s.elems, e)
		return true
	}
	return false
}

// Iter returns a channel of elements that can be ranged over.
func (s *Set) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		s.RLock()
		defer s.RUnlock()
		for elem := range s.elems {
			ch <- elem
		}
		close(ch)
	}()
	return ch
}

// ToSlice returns a slice containing all elements of this set.
func (s *Set) ToSlice() []interface{} {
	slice := make([]interface{}, 0, s.Size())
	s.RLock()
	defer s.RUnlock()
	for elem := range s.elems {
		slice = append(slice, elem)
	}
	return slice
}
