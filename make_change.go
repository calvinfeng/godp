package godp

// Coins is an array of coin values.
type Coins []int

// ForEach is a helper function to iterate through the list of coins.
func (c Coins) ForEach(f func(int)) {
	for _, denom := range c {
		f(denom)
	}
}

// DPFewestCoinsMakeChange is the dynamic programming version of finding minimum coin needed to make
// change for a sum.
func DPFewestCoinsMakeChange(c Coins, sum int) int {
	states := make([]int, sum+1)

	for i := 1; i < len(states); i++ {
		countCountChoices := []int{}
		c.ForEach(func(coin int) {
			remain := i - coin
			if remain >= 0 && states[remain] != -1 {
				countCountChoices = append(countCountChoices, states[remain]+1)
			}
		})

		states[i] = min(countCountChoices)
	}

	return states[len(states)-1]
}

func min(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}

	return min
}

// DPWaysToMakeChange computes the number of ways to make change for a given sum with a set of coin
// denominations. We perform memorization by recording and updating a state for each subproblem. The
// subproblem can be seen as incrementally making change for a subsum using 1 coin, 2 coin, and so
// forth...
func DPWaysToMakeChange(c Coins, sum int) int {
	states := make([]int, sum+1)

	states[0] = 1
	c.ForEach(func(coin int) {
		for i := 1; i < len(states); i++ {
			remain := i - coin
			if remain >= 0 {
				states[i] += states[remain]
			}
		}
	})

	return states[len(states)-1]
}
