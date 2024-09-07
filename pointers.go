package main

import "fmt"

func main() {
	i,j := 10,20
	p,q := &i,&j
	fmt.Println(p) //0xc000104040
	fmt.Println(*p) //10
	*p = 100
	fmt.Println(p) //0xc000104040
	fmt.Println(*p) //100
	m:= &p
	fmt.Println(m) //0xc000108038
	fmt.Println(q) //0xc000104048
}
