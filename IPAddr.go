package main

import "fmt"

type IPAddr [4]byte

func main() {
address := map[string]IPAddr {
"loopback":{127,0,0,1},
"Google DNS" :{8,8,8,8}
}

for v,a=range address {
fmt.Printf("%v : %v \n" , v, a)
}
}
