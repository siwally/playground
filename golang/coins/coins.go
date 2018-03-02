// Package coins provides functionality to calculate the ways
// of dividing a given number of coins into piles.
package coins

// PermsFn returns a function that can be called multiple times, in order to calculate
// the ways of dividing coins into piles.  Each time the function is called it increments
// the number of coins, starting from 1 coin.
func PermsFn() func() int {

	perms := make(map[int][][]int)
	coins := 0

	return func() int {
		coins++
		perms[coins] = [][]int{[]int{coins}} // add perm for all coins in one pile

		for i := 1; i < coins; i++ { // work upwards with a pile of 1, then 2, ...

			for _, p := range perms[coins-i] { // get perms for coins not in this pile

				if isNewPerm(p, i) {
					newPerms := appendPerm(p, i)
					perms[coins] = append(perms[coins], newPerms)
				}
			}
		}

		return len(perms[coins])
	}
}

func isNewPerm(perms []int, minPileSize int) bool {
	// walk from the end, as lowest values will be to the right
	for j := len(perms) - 1; j >= 0; j-- {
		if perms[j] < minPileSize {
			return false
		}
	}

	return true
}

// important to make a new array and not mutate existing perms
func appendPerm(perms []int, i int) []int {
	newPerms := make([]int, len(perms)+1)
	copy(newPerms, perms)
	newPerms[len(newPerms)-1] = i

	return newPerms
}
