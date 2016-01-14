package main

import "golang/x/tour/pic"

func Pic(dx , dy) [][]uint8 {
res := make([][]uint8, dy)
for y := range res{
    res[y]=make([]uint8, dx)
    for x := range res[y] {
     res[y][x]=uint8(x*y)
    }
    }
return res
}

func main() {
  pic.Show(Pic)
}
