package main

import (
	"fmt"

)

func main() {

	a:=A{
		100,
		0,
		10,
	}
	b:=B{
		300,
		0,
		20,
	}
	var i int
	var j int

	for i=1;i<=5;i++{
		a.blood=a.blood-b.agressivity
		b.blueblood=a.bluelood+10
		if b.blueblood == 50 {
			a.aggressivity=a.aggressivity-1
			b.blueblood=0
		}else {
			a.aggressivity=10
		}


	}
	fmt.Println(a.blood,b.blood)
	for j=1;j<=5;j++{
		b.blood=b.blood-a.aggressivity
	}
	fmt.Println(a.blood,b.blood)
}


type A struct {
	blood int
	bluelood int
	aggressivity int
}
type B struct {
	blood int
	blueblood int
	agressivity int
}

//func p (){
//	var m,n int
//	prpbabilities :=[]int{1,2}
//	rand.Intn(3)
//	for _,prpbabilities:=range prpbabilities{
//		m +=prpbabilities
//		if n<=rand &&m>rand{
//			return
//		}
//		n=m
//	}
//	return
