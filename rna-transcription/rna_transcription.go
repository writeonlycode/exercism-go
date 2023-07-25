// Package strand provides a function to determine the RNA complement of a
// given DNA sequence.
package strand

var complements = map[rune]string{'G': "C", 'C': "G", 'T': "A", 'A': "U"}

// ToRNA takes a DNA sequence as input and returns its RNA complement.
func ToRNA(dna string) string {
	result := ""

	for _, v := range dna {
		result += complements[v]
	}

	return result
}
