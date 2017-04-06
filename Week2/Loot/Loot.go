package Loot

import (
	"math"
)


// Something must go wrong, this thing can't be this complicated
func GetMaximumLoot (capacity int, values []int, weights []int) float64 {
	var units []float64
	for i := 0; i < len(values); i++ {
		units = append(units, float64(values[i])/float64(weights[i]))
	}

	var temp float64
	var tempWeights int
	for k := 0; k < len(units) - 2; k++ {
		for i := 0; i < len(units) - 1; i++ {
			if units[i] > units[i + 1] {
				temp = units[i]
				tempWeights = weights[i]

				units[i] = units[i + 1]
				weights[i] = weights[i + 1]

				units[i + 1] = temp
				weights[i + 1] = tempWeights
			}
		}
	}

	var result float64
	result = 0.0000

	for m := len(units) - 1; m >= 0; m-- {
		if weights[m] <= capacity {
			result += units[m] * float64(weights[m])
			capacity -= weights[m]
			if capacity == 0 {
				break
			} else {
				continue
			}
		}
		result += units[m] * float64(capacity)
		capacity = 0
		break
	}

	return RoundUp(result, 4)
}


func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}