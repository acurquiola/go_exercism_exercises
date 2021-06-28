package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	counter := 0
	if a == b {
		return counter, nil
	} else if len(a) != len(b) {
		return counter, errors.New("different lengths")
	} else if a != b {
		for i, _ := range a {
			if string(a[i]) != string(b[i]) {
				counter++
			}
		}
	}

	return counter, nil

}
