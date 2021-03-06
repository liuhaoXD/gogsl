package test

import (
	"fmt"
	"github.com/liuhaoXD/gogsl"
	"github.com/liuhaoXD/gogsl/sf"
	bessel "github.com/liuhaoXD/gogsl/sf/bessel/cylindrical"
	"math"
	"testing"
)

func TestSf(t *testing.T) {
	var x float64 = 5.0
	var result sf.GslSfResult
	expected := -0.17759677131433830434739701
	status := bessel.J0E(x, &result)
	fmt.Printf("status  = %s\n", gogsl.GslError(status).String())
	fmt.Printf("J0(5.0) = %.18f\n"+
		"      +/- % .18f\n", result.Val(), result.Err())
	fmt.Printf("exact   = %.18f\n", expected)
	if math.Abs(result.Val()-expected) > result.Err() {
		t.Error("incorrect value")
	}
}
