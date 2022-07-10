package square

import "errors"

var ErrorSize = errors.New("negative size")

func Square(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, ErrorSize
	}
	return a * b, nil
}
