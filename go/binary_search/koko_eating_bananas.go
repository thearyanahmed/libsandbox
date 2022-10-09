package binary_search

import (
	"math"
)

func KokoEatingBananas(piles []int, h int) int {

	l, r := 1, max(piles)

	res := r

	for l <= r {
		k := (l + r) / 2
		hours := 0

		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}

		if hours <= h {
			if k < res {
				res = k
			}

			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}

func max(v []int) int {
	max := v[0]

	for _, i := range v {
		if i > max {
			max = i
		}
	}

	return max
}
