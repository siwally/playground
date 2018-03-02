// Package sort provides functionality to sort integers, using an easy split /
// hard merge algorithm (classic von Neumann Mergesort).  The algorithm should
// yield O(n log n) performance, with the main overhead coming from walking
// back up the tree and merging pairs of sorted lists together.
package sort

// Interface alg allows the algorithm that orchestrates dividing the list into
// halves, sorting and merging to be changed in benchmarking, e.g. to use parallelism with channels.
type alg interface {
	divideAndMerge(in []int) []int
}

// Struct seqAlg is used as a method receiver for sequential orchestration of the algorithm
type seqAlg struct {
}

var algInUse alg = seqAlg{}

func Sort(in []int) ([]int, error) {
	return sort(in), nil // TODO Handle panics in defer block and return as errors
}

func sort(in []int) []int {
	var res []int

	switch len(in) {
	case 0, 1:
		res = in
	case 2:
		res = in

		if in[0] > in[1] {
			res[0], res[1] = in[1], in[0]
		}
	default:
		res = algInUse.divideAndMerge(in)
	}

	return res
}

func (s seqAlg) divideAndMerge(in []int) []int {
	l := in[0 : len(in)/2]
	r := in[len(l):]

	res := make([]int, len(in))
	merge(sort(l), sort(r), res)

	return res
}

// Too expensive for merge to make a new slice for the result each time it
// recurses, so pass as parameter res
func merge(l []int, r []int, res []int) {

	switch {
	case len(l) == 0:
		copy(res, r)
	case len(r) == 0:
		copy(res, l)
	case l[0] < r[0]:
		res[0], l = l[0], l[1:]
		merge(l, r, res[1:])
	default:
		res[0], r = r[0], r[1:]
		merge(l, r, res[1:])
	}
}
