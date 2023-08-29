package main
import (
	"fmt"
	"bytes"
)

func main(){
	sl1:=[]byte{'P','R','I','T'}
	sl2:=[]byte{'P','R','E','T'}

	res:=bytes.Compare(sl1,sl2)
	if res == 0{
		fmt.Println("Equal")
	}else{
		fmt.Println("Unequal")
	}
}