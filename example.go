package main

import (
"fmt"
)

func main(){

var x int
x=5
y:=10
z := x+y
name := "kjbcljascascbadhcidyg"
w:= max(x,y)

fmt.Println(x,"+",y,"=",z)
fmt.Println("your name is: ",name)
fmt.Println("max is:",w)

}
func max(a int, b int) int{
if a>b {
	return a
}else {
	return b
}
}
