package main

import "math"

func twoCrystalBalls(breaks []bool) int {
	jumps := int(math.Sqrt(float64((len(breaks)))))
	i := jumps

	for ; i < len(breaks); i += jumps {
		if breaks[i] {
			break
		}
	}

	i -= jumps

	for j := 0; j <= jumps && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		
		i++
	}

	return -1
}
