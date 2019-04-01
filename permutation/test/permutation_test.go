package test

import (
	"fmt"
	"github.com/liuhaoXD/gogsl/permutation"
	"github.com/liuhaoXD/gogsl/randist"
	"github.com/liuhaoXD/gogsl/rng"
	"os"
	"testing"
)

func TestPermutation(t *testing.T) {
	var N int = 15
	p := permutation.PermutationAlloc(N)
	q := permutation.PermutationAlloc(N)
	rng.EnvSetup()
	T := rng.DefaultRngType()
	r := rng.RngAlloc(T)
	fmt.Printf("initial permutation: ")
	permutation.PermutationInit(p)
	permutation.Fprintf(os.Stdout, p, " %u")
	os.Stdout.Sync()
	fmt.Printf("\n")
	fmt.Printf(" random permutation: ")
	randist.Shuffle(r, p.Slice_(), p.Len())
	permutation.Fprintf(os.Stdout, p, " %u")
	fmt.Printf("\n")
	fmt.Printf(" inverse permutation: ")
	permutation.Inverse(q, p)
	permutation.Fprintf(os.Stdout, q, " %u")
	fmt.Printf("\n")
}
