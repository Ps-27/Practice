package main
import "fmt"

func main(){

	fmt.Println("Reverse of number:")
    var n int
	fmt.Scan(&n)
	s:=0
	for ;n>0;{
		r:=n%10
		s=s*10+r
		n=n/10
	}
	fmt.Println("Reverse:",s)
}