// Package set implements a map-based set data structure.
package set

import (
	"fmt"
	"sync"
)

// Set represents the set data structure.
type Set struct {
	lock sync.RWMutex
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
	s.lock.Lock()
	defer s.lock.Unlock()
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
	s.lock.RLock()
	defer s.lock.RUnlock()
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
	s.lock.RLock()
	o.lock.RLock()
	defer s.lock.RUnlock()
	defer o.lock.RUnlock()
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
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elems = make(map[interface{}]struct{})
}

// IsEmpty returns true if this set contains no elements; false otherwise.
func (s *Set) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.size() == 0
}

// Size returns the total number of elements in this set.
func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.size()
}

func (s *Set) size() int {
	return len(s.elems)
}

// Remove removes each of the specified elements from this set, if present.
// Returns true if this set changed as a result of the call
func (s *Set) Remove(l ...interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
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
		s.lock.RLock()
		defer s.lock.RUnlock()
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
	s.lock.RLock()
	defer s.lock.RUnlock()
	for elem := range s.elems {
		slice = append(slice, elem)
	}
	return slice
}

// Subset determines if every item in the other set is in this set.
func (s *Set) Subset(o *Set) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for elem := range s.elems {
		if !o.contains(elem) {
			return false
		}
	}
	return true
}

// Superset determines if every item of this set is in the other set.
func (s *Set) Superset(o *Set) bool {
	return o.Subset(s)
}

// Union returns a new set with all items in both sets.
func (s *Set) Union(o *Set) *Set {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.union(o)
}

func (s *Set) union(o *Set) *Set {
	union := New()
	for elem := range o.elems {
		union.add(elem)
	}
	for elem := range s.elems {
		union.add(elem)
	}
	return union
}

// Intersect returns a new set with items that exist only in both sets.
func (s *Set) Intersect(o *Set) *Set {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.intersect(o)
}

func (s *Set) intersect(o *Set) *Set {
	intersect := New()
	// loop over smaller set
	if s.size() < o.size() {
		for elem := range s.elems {
			if o.contains(elem) {
				intersect.add(elem)
			}
		}
	} else {
		for elem := range o.elems {
			if s.contains(elem) {
				intersect.add(elem)
			}
		}
	}
	return intersect
}

// Diff returns a new set with items in the current set but not in the other
// set.
func (s *Set) Diff(o *Set) *Set {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.diff(o)
}

func (s *Set) diff(o *Set) *Set {
	diff := New()
	for elem := range s.elems {
		if !o.contains(elem) {
			diff.add(elem)
		}
	}
	return diff
}

// SymDiff returns a new set with items in the current set or the other set but
// not in both.
func (s *Set) SymDiff(o *Set) *Set {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.diff(o).union(o.diff(s))
}
