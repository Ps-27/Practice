package main
import (
	"fmt"
	"os"
	
)


func Calfactorial(number int )int{
	if number < 0{
		fmt.Println("invalid number")
		os.Exit(1)
	}
	if (number == 0){
		return 1
	}
	// fmt.Println("number:",number)
	return (number) * Calfactorial(number-1)
	
    
}
func main(){
	var number int
	fmt.Println("Enter number:")
	fmt.Scan(&number)
	factorial:=Calfactorial(number)
	fmt.Printf("Factorial of %v is %v.",factorial,factorial)

}