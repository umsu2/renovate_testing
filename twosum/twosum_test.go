package twosum_test

import (
	"errors"
	"testing"

	"github.com/umsu2/renovate_testing/twosum"
)

// testCase describes a single TwoSum scenario.
type testCase struct {
	name    string
	nums    []string
	target  string
	want    [2]int
	wantErr error
}

var cases = []testCase{
	// ── Basic integer cases ──────────────────────────────────────────────────
	{
		name:   "classic leetcode example 1",
		nums:   []string{"2", "7", "11", "15"},
		target: "9",
		want:   [2]int{0, 1},
	},
	{
		name:   "classic leetcode example 2",
		nums:   []string{"3", "2", "4"},
		target: "6",
		want:   [2]int{1, 2},
	},
	{
		name:   "same value twice",
		nums:   []string{"3", "3"},
		target: "6",
		want:   [2]int{0, 1},
	},
	{
		name:   "negatives sum to zero",
		nums:   []string{"-3", "0", "3"},
		target: "0",
		want:   [2]int{0, 2},
	},
	{
		name:   "large integers",
		nums:   []string{"1000000000", "2000000000", "3000000000"},
		target: "5000000000",
		want:   [2]int{1, 2},
	},

	// ── Decimal / fractional cases ───────────────────────────────────────────
	{
		name:   "simple decimals",
		nums:   []string{"1.5", "2.5", "3.0"},
		target: "4.0",
		want:   [2]int{0, 1},
	},
	{
		name:   "high-precision decimals that would fail with float64",
		nums:   []string{"0.1", "0.2", "0.7"},
		target: "0.3",
		// 0.1 + 0.2 = 0.3 — exact in decimal, imprecise in float64
		want: [2]int{0, 1},
	},
	{
		name:   "negative decimals",
		nums:   []string{"-1.5", "3.5", "2.0"},
		target: "2.0",
		want:   [2]int{0, 1},
	},
	{
		name:   "many decimal places",
		nums:   []string{"0.123456789", "0.987654321", "1.0"},
		target: "1.111111110",
		want:   [2]int{0, 1},
	},

	// ── Edge / error cases ───────────────────────────────────────────────────
	{
		name:    "no solution",
		nums:    []string{"1", "2", "3"},
		target:  "10",
		wantErr: twosum.ErrNoSolution,
	},
	{
		name:    "empty slice",
		nums:    []string{},
		target:  "5",
		wantErr: twosum.ErrNoSolution,
	},
	{
		name:    "single element",
		nums:    []string{"5"},
		target:  "5",
		wantErr: twosum.ErrNoSolution,
	},
	{
		name:    "invalid number in slice",
		nums:    []string{"1", "abc", "3"},
		target:  "4",
		wantErr: errors.New("invalid number"), // wrapped, check via errors.As indirectly
	},
	{
		name:    "invalid target",
		nums:    []string{"1", "2"},
		target:  "xyz",
		wantErr: errors.New("invalid target"),
	},
	{
		name:   "first and last elements",
		nums:   []string{"1", "5", "8", "9"},
		target: "10",
		want:   [2]int{0, 3},
	},
}

func TestTwoSum(t *testing.T) {
	t.Parallel()

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := twosum.TwoSum(tc.nums, tc.target)

			if tc.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr.Error())
				}
				// For sentinel errors use errors.Is; for descriptive errors
				// just verify the error is non-nil (checked above).
				if errors.Is(tc.wantErr, twosum.ErrNoSolution) &&
					!errors.Is(err, twosum.ErrNoSolution) {
					t.Fatalf("expected ErrNoSolution, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("TwoSum(%v, %q) = %v, want %v",
					tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

// BenchmarkTwoSum provides a basic performance baseline.
func BenchmarkTwoSum(b *testing.B) {
	nums := []string{"0.1", "0.2", "0.7", "1.5", "2.3"}
	target := "0.3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = twosum.TwoSum(nums, target)
	}
}
