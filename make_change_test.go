package godp

import (
	"testing"
)

type MakeChangeTestCase struct {
	Coins    Coins
	Sum      int
	Expected int
}

func TestDPFewestCoinMakeChange(t *testing.T) {
	testCases := []MakeChangeTestCase{
		MakeChangeTestCase{
			Coins:    Coins{1},
			Sum:      10,
			Expected: 10,
		},
		MakeChangeTestCase{
			Coins:    Coins{10},
			Sum:      10,
			Expected: 1,
		},
		MakeChangeTestCase{
			Coins:    Coins{12},
			Sum:      1,
			Expected: -1,
		},
		MakeChangeTestCase{
			Coins:    Coins{1, 3, 6},
			Sum:      13,
			Expected: 3,
		},
	}

	for _, tc := range testCases {
		actual := DPFewestCoinsMakeChange(tc.Coins, tc.Sum)
		if tc.Expected != actual {
			t.Errorf("expected %d but got %d", tc.Expected, actual)
		}
	}
}

func TestDPWaysToMakeChange(t *testing.T) {
	testCases := []MakeChangeTestCase{
		MakeChangeTestCase{
			Coins:    Coins{1},
			Sum:      10,
			Expected: 1,
		},
		MakeChangeTestCase{
			Coins:    Coins{1, 2, 5},
			Sum:      5,
			Expected: 4,
		},
		MakeChangeTestCase{
			Coins:    Coins{1, 2},
			Sum:      3,
			Expected: 2,
		},
		MakeChangeTestCase{
			Coins:    Coins{1, 2, 3},
			Sum:      7,
			Expected: 8,
		},
	}

	for _, tc := range testCases {
		actual := DPWaysToMakeChange(tc.Coins, tc.Sum)
		if tc.Expected != actual {
			t.Errorf("expected %d but got %d", tc.Expected, actual)
		}
	}
}
