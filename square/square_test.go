package square

import (
	"errors"
	"testing"
)

type testSquare struct {
	name string
	ina  int
	inb  int
	out  int
	err  error
}

func TestSquare(t *testing.T) {
	testSquares := []testSquare{
		{"1x2", 1, 2, 2, nil},
		{"2x3", 2, 3, 6, nil},
		{"3x4", 3, 4, 12, nil},
		{"4x5", 4, 5, 20, nil},
		{"5x6", 5, 6, 30, nil},
		{"6x7", 6, 7, 42, nil},
		{"-7x8", -7, 8, 56, ErrorSize},
		{"8x-9", 8, -9, 72, ErrorSize},
	}
	for _, sqr := range testSquares {
		sqr := sqr
		t.Run(sqr.name, func(t *testing.T) {
			result, err := Square(sqr.ina, sqr.inb)
			if err != nil {
				if errors.Is(err, sqr.err) {
					return
				}
				t.Error(err)
			}
			if result != sqr.out {
				t.Errorf("wtf %v", result)
			}
		})
	}

}
