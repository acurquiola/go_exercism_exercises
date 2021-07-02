package grains

import "errors"

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("n must be between 1 and 64 inclusive")
	}

	return uint64(0x1 << (input - 1)), nil

}

func Total() uint64 {
	var grainsSum uint64
	for i := 1; i <= 64; i++ {
		sum, _ := Square(i)
		grainsSum += sum
	}
	return grainsSum
}
