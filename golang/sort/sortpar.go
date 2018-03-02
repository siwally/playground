package sort

// Struct parAlg is used as a method receiver for parallel orchestration of the algorithm
type parAlg struct {
}

func (a parAlg) divideAndMerge(in []int) []int {

	l := in[0 : len(in)/2]
	r := in[len(l):]

	lout := make(chan []int)
	rout := make(chan []int)

	go func() {
		lout <- sort(l)
	}()

	go func() {
		rout <- sort(r)
	}()

	res := make([]int, len(l)+len(r))
	merge(<-lout, <-rout, res)

	return res
}
