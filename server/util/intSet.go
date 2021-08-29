package util

// IntSet is a collection of ints with no particular order
type IntSet map[int]struct{}

// Add adds any number of elements to the set
func (s *IntSet) Add(elements ...int) {
	for _, element := range elements {
		(*s)[element] = struct{}{}
	}
}

// Contains checks whether element is a member of the set
func (s *IntSet) Contains(element int) bool {
	_, included := (*s)[element]
	return included
}

// ContainsAll checks whether all elements are members of the set
func (s *IntSet) ContainsAll(elements ...int) bool {
	for _, element := range elements {
		if !s.Contains(element) {
			return false
		}
	}
	return true
}

// Contains checks whether at least one of the elements is a member of the set
func (s *IntSet) ContainsAny(elements ...int) bool {
	for _, element := range elements {
		if s.Contains(element) {
			return true
		}
	}
	return false
}

// Delete removes any number of elements from the set
func (s *IntSet) Delete(elements ...int) {
	for _, element := range elements {
		delete(*s, element)
	}
}

// Pop removes a random element of the set and returns it
func (s *IntSet) Pop() int {
	for element := range *s {
		s.Delete(element)
		return element
	}
	return -1
}

// Remove removes all elements of other from the set
func (s *IntSet) Remove(other IntSet) {
	for element := range other {
		s.Delete(element)
	}
}
