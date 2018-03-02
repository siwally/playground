package sort

type parAlg struct {
}

func (a parAlg) divideAndMerge(in []int) []int {

	l := in[0 : len(in)/2]
	r := in[len(l):]

	// return algs.merge(sort(l), sort(r)) // replace to use channel

	lout := make(chan []int)
	rout := make(chan []int)

	go func() {
		lout <- sort(l)
	}()

	go func() {
		rout <- sort(r)
	}()

	return merge(<-lout, <-rout)
}

// iterative merge

// func merge(l []int, r []int) []int {
// 	res := make([]int, len(l)+len(r))
//
// 	for i, j, x := 0, 0, 0; i < len(l) || j < len(r); x++ {
//
// 		if j == len(r) {
// 			copy(res[x:], l[i:])
// 			break
// 		}
//
// 		if i == len(l) {
// 			copy(res[x:], r[j:])
// 			break
// 		}
//
// 		if l[i] < r[j] {
// 			res[x] = l[i]
// 			i++
// 		} else {
// 			res[x] = r[j]
// 			j++
// 		}
// 	}
//
// 	return res
// }
