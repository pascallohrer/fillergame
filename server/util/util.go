package util

// FilterInt returns a set of all indices in slice who have the filter value, using type int
func FilterInt(slice []int, value int) IntSet {
	set := IntSet{}
	for key, element := range slice {
		if element == value {
			set.Add(key)
		}
	}
	return set
}

// Any uses test on each element of slice and returns true if at least one element passes the test
func Any(test func(interface{}) bool, values ...interface{}) bool {
	for _, element := range values {
		if test(element) {
			return true
		}
	}
	return false
}

// All uses test on each element of slice and returns true if every element passes the test
func All(test func(interface{}) bool, values ...interface{}) bool {
	for _, element := range values {
		if !test(element) {
			return false
		}
	}
	return true
}
