package main

import (
	"fmt"
	"sort"
)
// type Interface interface{

//     Len()int

//     Less(i, j int)bool

//     Swap(i,jint)
// }
type Human struct{
    name string
    age int
}
type AgeFactor []Human

func(a AgeFactor )Len()int {return len(a)}
func(a AgeFactor)Less(i,j int)bool{return a[i].age<a[j].age}
func(a AgeFactor)Swap(i,j int){ a[i],a[j]=a[j],a[i]}

func main(){

    person:=[]Human{
        {"Sachin",30},
        {"Tanu",13},
        {"manu",20},
    }
    sort.Sort(AgeFactor(person))
    fmt.Println(person)
}