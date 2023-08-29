package main
import (
	"fmt"
	"strings"
)


func swapContent(obj []string){
	for i,j:=0,len(obj)-1;i<j;i,j=i+1,j-1{
		obj[i], obj[j]=obj[j],obj[i]
	}
}
func main(){
	ob:=[]string{"prity","sinha"}
	swapContent(ob)
	fmt.Println(ob)



}