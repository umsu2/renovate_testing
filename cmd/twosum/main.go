// Command twosum is a CLI wrapper around the twosum package.
// Usage: twosum <target> <num1> <num2> [num3 ...]
// Example: twosum 9 2 7 11 15
package main

import (
	"fmt"
	"os"

	"github.com/umsu2/renovate_testing/twosum"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "usage: twosum <target> <num1> <num2> [num3 ...]")
		os.Exit(1)
	}

	target := os.Args[1]
	nums := os.Args[2:]

	indices, err := twosum.TwoSum(nums, target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("indices: [%d, %d]  (%s + %s = %s)\n",
		indices[0], indices[1],
		nums[indices[0]], nums[indices[1]], target)
}
