// Package strain defines filter methods for slices of ints, list of ints, and
// list of strings.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep filters out the Ints that don't satisfy the given predicate.
func (i Ints) Keep(filter func(int) bool) Ints {
	if len(i) == 0 {
		return nil
	} else {
		if filter(i[len(i)-1]) {
			return append(i[0:(len(i)-1)].Keep(filter), i[len(i)-1])
		} else {
			return i[0:(len(i) - 1)].Keep(filter)
		}
	}
}

// Discard filters out the Ints that do satisfy the given predicate.
func (i Ints) Discard(filter func(int) bool) Ints {
	if len(i) == 0 {
		return nil
	} else {
		if !filter(i[len(i)-1]) {
			return append(i[0:(len(i)-1)].Discard(filter), i[len(i)-1])
		} else {
			return i[0:(len(i) - 1)].Discard(filter)
		}
	}
}

// Keep filters out the Lists that don't satisfy the given predicate.
func (l Lists) Keep(filter func([]int) bool) Lists {
	if len(l) == 0 {
		return nil
	} else {
		if filter(l[len(l)-1]) {
			return append(l[0:(len(l)-1)].Keep(filter), l[len(l)-1])
		} else {
			return l[0:(len(l) - 1)].Keep(filter)
		}
	}
}

// Keep filters out the Strings that don't satisfy the given predicate.
func (s Strings) Keep(filter func(string) bool) Strings {
	if len(s) == 0 {
		return nil
	} else {
		if filter(s[len(s)-1]) {
			return append(s[0:(len(s)-1)].Keep(filter), s[len(s)-1])
		} else {
			return s[0:(len(s) - 1)].Keep(filter)
		}
	}
}
