package main
import "fmt"


func reverse(str string)string{
	bs:=[]byte(str)
	fmt.Println(bs)
	rev:=make([]byte,len(bs))
	for i,v:=range bs{
		rev[len(bs)-i-1]=v
	}
	fmt.Println("rev bs:",rev)
	return string(rev)
}

func ReverseString(str string)(result string){

	for _, v:=range str{
		result=string(v)+result
	}
	
	return 
}
func main(){

	fmt.Print("Enter String:")
    var s string
	fmt.Scan(&s)

	fmt.Println("original string:",s)
	fmt.Println("Reverse string:",reverse(s))
	fmt.Println("original string:",s)
	fmt.Println("Reverse string:",ReverseString(s))

}