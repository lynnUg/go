package main

import (
"fmt"
"math"
)

const DELTA=0.0001
const START_Z=100

func Sqrt( x float64) float64 {
z := START_z
step := func (z float64) float64 {
return z-(z*z-x)/(2*z)
}
for zz:=step();math.Abs(zz-z)>DELTA {
z=zz
zz=step()
}
}

func main() {
fmt.Println(Sqrt(500))
fmt.Println(math.Sqrt(500))

}
