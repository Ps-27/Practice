package main
import "fmt"


func findNficbonacci(n int)int{
	FirstNumber,SecNumber:=0,1
	fmt.Println(FirstNumber," ",SecNumber)
	var nextNumber int
	for c:=2;c<=n;c++{
		nextNumber = FirstNumber + SecNumber
		FirstNumber=SecNumber
		SecNumber=nextNumber
		fmt.Print(" ",nextNumber)
	}
	return nextNumber
	
}


func fibo(n int )int{
	if n<2{
		return n
	}
	return fibo(n-1)+fibo(n-2)
}
func main(){
	var  number int
	fmt.Println("Enter nth number:")
	fmt.Scan(&number)
	fmt.Println("nth fibonacci number:",findNficbonacci(number))
	fmt.Println("fibo:",fibo(number))

}