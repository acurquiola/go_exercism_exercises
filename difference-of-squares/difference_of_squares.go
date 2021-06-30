package diffsquares

import (
	"math"
)

func SquareOfSum(N int) int {
	sqOfSum := 0

	for i := 0; i < N; i++ {
		sqOfSum += i + 1
	}

	sqOfSum = int(math.Pow(float64(sqOfSum), 2))

	return sqOfSum
}

func SumOfSquares(N int) int {
	sumOfSquares := 0

	for i := 0; i < N; i++ {
		sumOfSquares += int(math.Pow(float64(i+1), 2))
	}

	return sumOfSquares
}

func Difference(N int) int {
	return int(math.Abs(float64(SumOfSquares(N) - SquareOfSum(N))))
}
