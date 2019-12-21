package main

import (
	"fmt"
	"io/ioutil"
	"os"

)


func main() {

	file, err := os.Open("Student.txt")
	if err != nil{
		fmt.Printf("open file err=%v \n",err)
		return
	}
	defer file.Close()

	go Name("573")

}
func Name(aname string)(n int){
	date,_:=ioutil.ReadFile("Student.txt")
	Name:=string(date)
	namechange:=([]rune)(Name)
	name:=make([]string,1,1)
	t:=""
	for x:=range namechange{
		if string(namechange[x])!="\n"{
			t=t+string(namechange[x])
		}else{
			name=append(name,t)
			t=""
		}
	}
	for x:=0;x<len(name);x++{
		if name[x]==aname{
			n++
		}
	}
	return n
}