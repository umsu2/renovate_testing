// Package twosum provides a Two Sum solution using precise decimal arithmetic.
// It uses github.com/shopspring/decimal to avoid floating-point inaccuracies.
package twosum

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// ErrNoSolution is returned when no two indices sum to the target.
var ErrNoSolution = errors.New("no two-sum solution found")

// TwoSum finds the indices of two numbers in nums that add up to target.
// All arithmetic uses shopspring/decimal for exact decimal calculation.
// Returns the pair of indices [i, j] where i < j, or an error if none exist.
func TwoSum(nums []string, target string) ([2]int, error) {
	t, err := decimal.NewFromString(target)
	if err != nil {
		return [2]int{}, fmt.Errorf("invalid target %q: %w", target, err)
	}

	// Map from complement (as string) to its index.
	seen := make(map[string]int)

	for i, raw := range nums {
		n, err := decimal.NewFromString(raw)
		if err != nil {
			return [2]int{}, fmt.Errorf("invalid number at index %d %q: %w", i, raw, err)
		}

		complement := t.Sub(n)
		key := complement.String()

		if j, ok := seen[key]; ok {
			return [2]int{j, i}, nil
		}

		seen[n.String()] = i
	}

	return [2]int{}, ErrNoSolution
}
