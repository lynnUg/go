package main

import "fmt"

func fibonnaci() func() int {
x, y := 0, 1
return func() int {
x,y =x, y+x

return y-x
}
}

func main() {
f := fibonnaci()

for i :=0 ;i <10 ;i++{
fmt.Println(f(i))
}
}
