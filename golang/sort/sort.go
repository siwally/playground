// Package sort provides functionality to sort integers, using an easy split /
// hard merge algorithm (classic von Neumann Mergesort).  The algorithm should
// yield O(n log n) performance, with the main overhead coming from walking
// back up the tree and merging pairs of sorted lists together.
package sort

type alg interface {
	divideAndMerge(in []int) []int
}

type seqAlg struct {
}

var a alg = seqAlg{}

func Sort(in []int) ([]int, error) {
	return sort(in), nil // TODO Handle panics in defer block and return as errors
	// TODO Try allocating single array to use with recursion
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
		res = a.divideAndMerge(in)
	}

	return res
}

func (s seqAlg) divideAndMerge(in []int) []int {
	l := in[0 : len(in)/2]
	r := in[len(l):]

	return merge(sort(l), sort(r))
}

func merge(l []int, r []int) []int {

	res := make([]int, len(l)+len(r))

	switch {
	case len(l) == 0:
		return r
	case len(r) == 0:
		return l
	case l[0] < r[0]:
		res[0], l = l[0], l[1:]
	default:
		res[0], r = r[0], r[1:]
	}

	copy(res[1:], merge(l, r))
	return res
}
