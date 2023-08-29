package main
import (
	"fmt"
	"strings"
)

func main(){
	var n  int
	FisrtSmallNum,SecondSmallNum:=0,0
	fmt.Scan(&n)
	a:=make([]int,n)
	for i,_:=range a{
		fmt.Scan(&a[i])
	}
	if a[0]<a[1]{
		FisrtSmallNum=a[0]
		SecondSmallNum=a[1]
	}else{
		FisrtSmallNum=a[1]
		SecondSmallNum=a[0]
	}
	for i,_:=range a{
		fmt.Println("ith ele:",i,"a[i]",a[i])
		if a[i]< FisrtSmallNum{
			SecondSmallNum=FisrtSmallNum
			FisrtSmallNum=a[i]
		}else if a[i]< SecondSmallNum  && a[i]>FisrtSmallNum{
			SecondSmallNum=a[i]
		}
	}

	fmt.Println(FisrtSmallNum)
	fmt.Println(SecondSmallNum)

	s:="prity"
	ss:="sinha"
	sss:=s+ss
	slice:=[]string{"prity","sinha"}
	sslice:=strings.Join(slice,"-")
	fmt.Println(sslice)
	fmt.Println(sss)
	count:=make(map[rune]int)
	for _,char:=range sss{
		count[char]++
	}
    c:=make(map[string]int)
	for i, v:=range count{
		c[string(i)]=v
	}
	for i,_:=range count{
		fmt.Printf("%c occurs %d times\n",i,count[i])
	}
	fmt.Println(count)
	for i:=0;i<len(sslice);i++{
		for j:=i+1;j<len(sslice);j++{
			if sslice[i] == sslice[j]{
				fmt.Printf("%c occurs twice\n",sslice[j])
				break
			}
		}
	}
	fmt.Println(c)
	fmt.Println("Done")
}
