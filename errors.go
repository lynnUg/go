package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64
const DELTA=0000.1
const INTIAL_Z=10.0
func (e ErrNegativeSqrt) Error() string {
return fmt.Sprintf("cannot Sqrt negative number: %v",e)
}

func Sqrt(x float64) (float64, error) {
	if x<0.0 {
	return 0, ErrNegativeSqrt(x)
	}
	z :=INTIAL_Z
	step := func () float64 {
	return z-(z*z -x)/(2*x)
	}
	
	for zz:=step() ; math.Abs(zz-z)>DELTA
	{
	 z =zz
	 zz =step()
	}
	
	return z, nil
	
}

func main() {
	//fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
